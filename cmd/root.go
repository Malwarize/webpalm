package cmd

import (
	"fmt"
	"github.com/XORbit01/webpalm/core"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "webpalm",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			//help message
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("URL: ", url)
		level, err := cmd.Flags().GetInt("level")
		if err != nil {
			//help message
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Level: ", level)

		liveMode, err := cmd.Flags().GetBool("live")
		if err != nil {
			//help message
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Is live mode: ", liveMode)

		if level < 1 {
			fmt.Println("Error: Level should be greater than 0")
			return
		}
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			fmt.Println("you didn't specify the protocol, so we will use http")
			url = "http://" + url
		}
		exportFile, err := cmd.Flags().GetString("output")
		fmt.Println("Export File: ", exportFile)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		regexMap, err := cmd.Flags().GetStringToString("regexes")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Regex List: ", regexMap)

		excludedStatus, err := cmd.Flags().GetIntSlice("exclude-code")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Status Code List: ", excludedStatus)
		cr := core.NewCrawler(url, level, liveMode, exportFile, regexMap, excludedStatus)
		cr.Crawl()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "URL to the website")
	err := rootCmd.MarkFlagRequired("url")
	if err != nil {
		return
	}
	rootCmd.Flags().IntP("level", "l", 1, "Level of the website to crawl")

	rootCmd.Flags().Bool("live", false, "Live output mode")

	rootCmd.Flags().StringP("output", "o", "", "Output file name")

	rootCmd.Flags().StringToString("regexes", map[string]string{}, "Regexes to match")

	rootCmd.Flags().IntSliceP("exclude-code", "x", []int{}, "Status codes to exclude")
}
