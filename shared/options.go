package shared

import (
	"fmt"
	"net"
	urlTool "net/url"
	"reflect"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type Options struct {
	URL   string `name:"url"`
	Level int    `name:"level"`
	//LiveMode        bool              `name:"live"`
	ExportFile      string            `name:"save to file"`
	RegexMap        map[string]string `name:"regexes"`
	StatusResponses []int             `name:"exclude codes"`
	IncludedUrls    []string          `name:"include"`
	Workers         int               `name:"workers"`
	Delay           int               `name:"delay"`
	Proxy           *urlTool.URL      `name:"proxy"`
	TimeOut         int               `name:"timeout"`
	UserAgent       string            `name:"user agent"`
}

func (o *Options) BuildOptionBanner() string {
	var banner string
	banner += color.RedString("┌")
	banner += color.RedString("[")
	banner += color.MagentaString((*o).URL)
	banner += color.RedString("]\n")
	t := reflect.TypeOf(*o)
	for i := 0; i < t.NumField(); i++ {

		field := t.Field(i)
		name := field.Tag.Get("name")
		if name == "url" {
			continue
		}
		value := reflect.ValueOf(*o).Field(i).Interface()
		if value == nil || value == "" || value == false {
			banner += color.RedString("│")
			banner += color.BlueString(name + ": ")
			banner += color.CyanString("not set")
			banner += "\n"
			continue
		}

		typeof := reflect.TypeOf(value)
		if typeof.Kind() == reflect.Slice {
			s := reflect.ValueOf(value)
			if s.Len() == 0 {
				banner += color.RedString("│")
				banner += color.BlueString(name + ": ")
				banner += color.CyanString("not set")
				banner += "\n"
				continue
			}
			banner += color.RedString("│")
			banner += color.BlueString(name + ": ")

			// check the type of the slice
			typeOfIndex := s.Index(0).Kind()
			if typeOfIndex == reflect.Int {
				for i := 0; i < s.Len(); i++ {
					if i == s.Len()-1 {
						banner += color.CyanString(" %d", s.Index(i).Interface().(int))
						continue
					}

					banner += color.CyanString(" %d", s.Index(i).Interface().(int))
					banner += color.CyanString(",")
				}
			} else if typeOfIndex == reflect.String {
				for i := 0; i < s.Len(); i++ {
					banner += color.CyanString("\n")
					banner += color.RedString("│")
					banner += color.CyanString("  %s", s.Index(i).Interface().(string))
				}
			}
			banner += "\n"

		} else if typeof.Kind() == reflect.String {

			banner += color.RedString("│")
			banner += color.BlueString(name + ": ")
			banner += color.CyanString("%s", value.(string))
			banner += "\n"
		} else if typeof.Kind() == reflect.Int {
			banner += color.RedString("│")
			banner += color.BlueString(name + ": ")
			banner += color.CyanString("%d", value.(int))
			banner += "\n"
		} else if typeof.Kind() == reflect.Bool {
			banner += color.RedString("│")
			banner += color.BlueString(name + ": ")
			banner += color.CyanString("%t", value.(bool))
			banner += "\n"
		} else if typeof.Kind() == reflect.Map {
			if len(value.(map[string]string)) == 0 {
				banner += color.RedString("│")
				banner += color.BlueString(name + ": ")
				banner += color.CyanString("not set")
				banner += "\n"
				continue
			}
			banner += color.RedString("│")
			banner += color.BlueString(name + ": ")
			banner += color.CyanString("\n")
			for k, v := range value.(map[string]string) {
				banner += color.RedString("│")
				banner += color.CyanString("  %s: %s", k, v)
				banner += "\n"
			}
		} else {

			banner += color.RedString("│")
			banner += color.BlueString(name + ": ")
			if fmt.Sprintf("%v", value) == "<nil>" {
				banner += color.CyanString("not set")
			} else {
				banner += color.CyanString("%v", value)
			}
			banner += "\n"
		}

	}
	banner += color.RedString("└")
	return banner
}

func (o *Options) PrintBanner() {
	fmt.Println(o.BuildOptionBanner())
}

func (o *Options) ManipulateData() {
	if matched, _ := regexp.MatchString(`https?://`, o.URL); !matched {
		o.URL = "http://" + o.URL
	}
}

func IsValidDomain(url string) bool {
	// check if url is an ip address
	if ip := net.ParseIP(url); ip != nil {
		return true
	}
	if regexp.MustCompile(`^([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$`).MatchString(url) {
		return true
	}
	return false
}

func ValidateThenBuildOption(cmd *cobra.Command) (*Options, error) {
	url, err := cmd.Flags().GetString("url")
	if err != nil {
		return nil, err
	}
	level, err := cmd.Flags().GetInt("level")
	if err != nil {
		return nil, err
	}

	if level < 0 {
		return nil, err
	}
	exportFile, err := cmd.Flags().GetString("output")
	if err != nil {
		return nil, err
	}
	regex, err := cmd.Flags().GetString("regexes")
	if err != nil {
		return nil, err
	}
	regexMap := make(map[string]string)
	ss := strings.Split(regex, `",`)
	lastkey := ""
        for _, v := range ss {
                r := strings.Split(v, `="`)
                regexMap[r[0]] = r[1]
		lastkey = r[0]
        }
	// handle the last edge case
	regexMap[lastkey] = regexMap[lastkey][:len(regexMap[lastkey])-1]

        for k, v := range regexMap {
                println(k, v)
        }
	
	excludedStatus, err := cmd.Flags().GetIntSlice("exclude-code")
	if err != nil {
		return nil, err
	}

	includedUrls, err := cmd.Flags().GetStringSlice("include")
	if err != nil {
		return nil, err
	}
	for _, include := range includedUrls {
		if !IsValidDomain(include) {
			return nil, fmt.Errorf("invalid domain  %s", include)
		}
	}
	workers, err := cmd.Flags().GetInt("worker")
	if err != nil {
		return nil, err
	}

	delay, err := cmd.Flags().GetInt("delay")
	if err != nil {
		return nil, err
	}

	proxy, err := cmd.Flags().GetString("proxy")
	if err != nil {
		return nil, err
	}

	var parsedProxy *urlTool.URL
	if proxy != "" {
		parsedProxy, err = urlTool.Parse(proxy)
		if err != nil {
			return nil, err
		}

	} else {
		parsedProxy = nil
	}

	timeout, err := cmd.Flags().GetInt("timeout")
	if err != nil {
		return nil, err
	}
	userAgent, err := cmd.Flags().GetString("user-agent")
	if err != nil {
		return nil, err
	}

	var userAgentString string
	userAgent = strings.ToLower(userAgent)
	if userAgent == "chrome" {
		userAgentString = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36" +
			"(KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36"
	} else if userAgent == "firefox" {
		userAgentString = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:63.0) Gecko/20100101" +
			" Firefox/63.0"
	} else if userAgent == "safari" {
		userAgentString = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14) AppleWebKit/605.1.15" +
			" (KHTML, like Gecko) Version/12.0 Safari/605.1.15"
	} else if userAgent == "ie" {
		userAgentString = "Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0)" +
			" like Gecko"
	} else if userAgent == "opera" {
		userAgentString = "Opera/9.80 (Windows NT 6.1; WOW64) Presto/2.12.388 Version/12.18"
	} else if userAgent == "android" {
		userAgentString = "Mozilla/5.0 (Linux; Android 8.0.0;) AppleWebKit/537.36" +
			" (KHTML, like Gecko) Chrome/70.0.3538.110 Mobile Safari/537.36"
	} else if userAgent == "ios" {
		userAgentString = "Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/605.1.15" +
			" (KHTML, like Gecko) Version/12.0 Mobile/15E148 Safari/604.1"
	} else {
		userAgentString = userAgent
	}
	if workers > 0 && delay > 0 {
		return nil, fmt.Errorf("you can't use delay in paralell mode")
	}
	options := &Options{
		URL:             url,
		Level:           level,
		ExportFile:      exportFile,
		RegexMap:        regexMap,
		StatusResponses: excludedStatus,
		IncludedUrls:    includedUrls,
		Workers:         workers,
		Delay:           delay,
		Proxy:           parsedProxy,
		TimeOut:         timeout,
		UserAgent:       userAgentString,
	}
	options.ManipulateData()
	return options, nil
}
