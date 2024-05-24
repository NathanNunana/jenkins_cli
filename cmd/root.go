package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	jenkinsURL string
	username   string
	apiToken   string
	path       string
)

var rootCmd = &cobra.Command{
	Use:   "jenkins-cli",
	Short: "A CLI tool to interact with jenkins",
	Long:  `A simple CLI tool built in Go to interact with jenkins`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func init() {
}
