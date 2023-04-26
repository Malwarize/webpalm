/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

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
}
