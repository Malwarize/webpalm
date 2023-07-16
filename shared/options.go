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
	URL             string            `name:"url"`
	Level           int               `name:"level"`
	LiveMode        bool              `name:"live"`
	ExportFile      string            `name:"save to file"`
	RegexMap        map[string]string `name:"regexes"`
	StatusResponses []int             `name:"exclude codes"`
	IncludedUrls    []string          `name:"include"`
	MaxConcurrency  int               `name:"max concurrency"`
	Proxy           *urlTool.URL      `name:"proxy"`
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
	if !strings.HasPrefix(o.URL, "http://") && !strings.HasPrefix(o.URL, "https://") {
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

	liveMode, err := cmd.Flags().GetBool("live")
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
	regexMap, err := cmd.Flags().GetStringToString("regexes")
	if err != nil {
		return nil, err
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
	maxConcurrency, err := cmd.Flags().GetInt("max-concurrency")
	if err != nil {
		return nil, err
	}
	if maxConcurrency < 1 {
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
	// set max concurrency to 1 if live mode is enabled
	if liveMode {
		maxConcurrency = 1
	}

	options := &Options{
		URL:             url,
		Level:           level,
		LiveMode:        liveMode,
		ExportFile:      exportFile,
		RegexMap:        regexMap,
		StatusResponses: excludedStatus,
		IncludedUrls:    includedUrls,
		MaxConcurrency:  maxConcurrency,
		Proxy:           parsedProxy,
	}
	options.ManipulateData()
	return options, nil
}
