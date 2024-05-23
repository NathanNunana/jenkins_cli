package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	// system username
	out, err := exec.Command("whoami").Output()
	if err != nil {
		log.Fatal(err)
	}

	user := strings.TrimSpace(string(out))
	path = fmt.Sprintf("/home/%s/.jenkins/", user)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if _, err := os.Stat(path); err == nil {
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				log.Fatalf("Error reading config file, %v", err)
			} else {
				log.Printf("another error, %v", err)
				log.Printf("%v", viper.Get("credentials"))
			}
			jenkinsURL = viper.GetString("credentials")
			fmt.Println(jenkinsURL)
		}
	}
}
