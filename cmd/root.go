/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
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
		if url == "" {
			fmt.Println("Error: URL is required")
			return
		}
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

		outputMode, err := cmd.Flags().GetString("mode-output")
		if err != nil {
			//help message
			fmt.Println("Error: ", err)
			return
		}
		if outputMode != "live" && outputMode != "block" {
			fmt.Println("Error: Output mode should be either live or block")
			return
		}
		fmt.Println("Output Mode: ", outputMode)

		if level < 1 {
			fmt.Println("Error: Level should be greater than 0")
			return
		}
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			fmt.Println("you didn't specify the protocol, so we will use http")
			url = "http://" + url
		}
		cr := core.NewCrawler(url, level, outputMode)
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

	rootCmd.Flags().IntP("level", "l", 1, "Level of the website to crawl")

	rootCmd.Flags().StringP("mode-output", "m", "block", "Output mode: live, block ")

}
