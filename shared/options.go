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
	var sb strings.Builder

	sb.WriteString(color.RedString("┌["))
	sb.WriteString(color.MagentaString((*o).URL))
	sb.WriteString(color.RedString("]\n"))

	t := reflect.TypeOf(*o)
	v := reflect.ValueOf(*o)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := field.Tag.Get("name")
		if name == "url" {
			continue
		}

		value := v.Field(i).Interface()
		if value == nil || value == "" || value == false {
			sb.WriteString(color.RedString("│"))
			sb.WriteString(color.BlueString(name + ": "))
			sb.WriteString(color.CyanString("not set\n"))
			continue
		}
		sb.WriteString(color.RedString("│"))
		sb.WriteString(color.BlueString(name + ": "))

		switch value := value.(type) {
		case []string:
			if len(value) == 0 {
				sb.WriteString(color.CyanString("not set"))
			} else {
				sb.WriteString(color.CyanString("\n"))
				for _, s := range value {
					sb.WriteString(color.RedString("│"))
					sb.WriteString(color.CyanString("  %s", s))
					sb.WriteString("\n")
				}
			}
		case []int:
			if len(value) == 0 {
				sb.WriteString(color.CyanString("not set"))
			} else {
				for i, v := range value {
					if i == 0 {
						sb.WriteString(color.CyanString("%d", v))
					} else {
						sb.WriteString(color.CyanString(", %d", v))
					}
				}
			}
		case map[string]string:
			if len(value) == 0 {
				sb.WriteString(color.CyanString("not set"))
			} else {
				sb.WriteString(color.CyanString("\n"))
				for k, v := range value {
					sb.WriteString(color.RedString("│"))
					sb.WriteString(color.CyanString("  %s: %s", k, v))
					sb.WriteString("\n")
				}
			}
		case string:
			sb.WriteString(color.CyanString("%s", value))
		case int:
			sb.WriteString(color.CyanString("%d", value))
		case bool:
			sb.WriteString(color.CyanString("%t", value))
		default:
			if fmt.Sprintf("%v", value) == "<nil>" {
				sb.WriteString(color.CyanString("not set"))
			} else {
				sb.WriteString(color.CyanString("%v", value))
			}
		}

		sb.WriteString("\n")
	}

	sb.WriteString(color.RedString("└"))

	return sb.String()
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
	// regexMap, err := cmd.Flags().GetStringToString("regexes")
	regexMap := cmd.Flags().Lookup("regexes").Value.(*RegexFlag).Value()

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
