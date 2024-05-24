package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Set up credentials to use jenkins server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// directory for credentials
		if _, err := os.Stat(path); os.IsNotExist(err) {
			if err := os.Mkdir(path, 0750); err != nil {
				log.Fatalf("Error creating dir, %v", err)
			}
		}
		// writing credentials
		filePath := fmt.Sprintf("%s/credentials", path)
		// data := fmt.Sprintf("credentials:\n\turl: %s\n\tusername: %s\n\tapiToken: %s\n", jenkinsURL, username, apiToken)
		data := fmt.Sprintf("url=%s\nusername=%s\napiToken=%s\n", jenkinsURL, username, apiToken) // required data
		if err := os.WriteFile(filePath, []byte(data), 0644); err != nil {
			log.Fatalf("Configuration failed, %v\n", err)
		}

		fmt.Println("Crendentials configured successfully")
	},
}

func init() {
	// configuration command, command flags
	rootCmd.AddCommand(configureCmd)
	configureCmd.Flags().StringVarP(&jenkinsURL, "url", "u", "", "jenkins URL")
	configureCmd.Flags().StringVarP(&username, "username", "n", "", "jenkins username")
	configureCmd.Flags().StringVarP(&apiToken, "token", "t", "", "jenkins api token")

	// required fields
	_ = configureCmd.MarkFlagRequired("url")
	_ = configureCmd.MarkFlagRequired("username")
	_ = configureCmd.MarkFlagRequired("token")
}
