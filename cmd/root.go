/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sauce",
	Short: "🥫 sauce is a CLI tool to find the source of anime screenshots, gifs, clips, etc.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.SetUsageTemplate(`Usage:
  sauce [command]

Available Commands:
  file        Search using a local file.
  url         Search for anime source using the url to the media.
  help        Help about any command.

Use "sauce [command] --help" for more information about a command.
`)
	rootCmd.AddCommand(urlCmd)
	rootCmd.AddCommand(fileCmd)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
