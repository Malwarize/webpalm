package cmd

import (
	"fmt"
	"github.com/Malwarize/webpalm/v2/core"
	"github.com/spf13/cobra"
	"net"
	"os"
	"regexp"
	"strings"
)

func isValidDomain(url string) bool {
	//check if url is an ip address
	if ip := net.ParseIP(url); ip != nil {
		return true
	}
	if regexp.MustCompile(`^([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$`).MatchString(url) {
		return true
	}
	return false
}

var rootCmd = &cobra.Command{
	Use:   usage(),
	Short: "A web scraping tool",
	Long:  long(),
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			//help message
			fmt.Println("Error: ", err)
			return
		}
		level, err := cmd.Flags().GetInt("level")
		if err != nil {
			//help message
			fmt.Println("Error: ", err)
			return
		}

		liveMode, err := cmd.Flags().GetBool("live")
		if err != nil {
			//help message
			fmt.Println("Error: ", err)
			return
		}

		if level < 0 {
			fmt.Println("Error: Level should be greater equal than 0")
			return
		}
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		exportFile, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		regexMap, err := cmd.Flags().GetStringToString("regexes")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		excludedStatus, err := cmd.Flags().GetIntSlice("exclude-code")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		includedUrls, err := cmd.Flags().GetStringSlice("include")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		for _, include := range includedUrls {
			if !isValidDomain(include) {
				fmt.Println("Error: Invalid domain name: ", include)
				return
			}
		}
		maxConcurrency, err := cmd.Flags().GetInt("max-concurrency")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		if maxConcurrency < 1 {
			fmt.Println("Error: Max concurrency should be greater equal than 1")
			return
		}
		fmt.Println(banner())
		fmt.Println(options(url, level, liveMode, exportFile, regexMap, excludedStatus, includedUrls, maxConcurrency))
		cr := core.NewCrawler(url, level, liveMode, exportFile, regexMap, excludedStatus, includedUrls, maxConcurrency)
		cr.Crawl()
	},
	Example: example() + regexestable(),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "target url / ex: -u https://google.com")
	if err := rootCmd.MarkFlagRequired("url"); err != nil {
		return
	}
	rootCmd.Flags().IntP("level", "l", 0, "level of palming / ex: -l2")

	rootCmd.Flags().Bool("live", false, "live output mode (slow but live streaming) use only 1 thread / ex: --live")

	rootCmd.Flags().StringP("output", "o", "", "file to export the result (f.json, f.xml, f.txt) / ex: -o result.json")

	rootCmd.Flags().StringToString("regexes", map[string]string{}, "regexes to match in each page / ex: --regexes comments=\"\\<\\!--.*?-->\"")

	rootCmd.Flags().IntSliceP("exclude-code", "x", []int{}, "status codes to exclude / ex : -x 404,500")

	rootCmd.Flags().StringSliceP("include", "i", []string{}, "include only domains / ex : -i google.com,facebook.com")

	rootCmd.Flags().IntP("max-concurrency", "m", 1000, "max concurrent tasks / ex: -m 10")
}
