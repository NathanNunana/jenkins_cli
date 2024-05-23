package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

type Job struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Color string `json:"color"`
}

type JobResponse struct {
	Jobs []Job `json:"jobs"`
}

var listJobCmd = &cobra.Command{
	Use:   "list",
	Short: "list all jenkins jobs",
	Long:  "Produces a list of all jenkins jobs created",
	Run: func(cmd *cobra.Command, args []string) {
		// api call
		url := fmt.Sprintf("%s/api/json", jenkinsURL)
		client := http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}

		// authenticate
		req.SetBasicAuth(username, apiToken)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Fatalf("Failed to load jobs, %v", resp.StatusCode)
		}

		var jenkinsResponse JobResponse

		// decode response object object
		if err := json.NewDecoder(resp.Body).Decode(&jenkinsResponse); err != nil {
			log.Fatalf("Failed to decode response, %v", err)
		}

		// display the response
		for _, job := range jenkinsResponse.Jobs {
			fmt.Printf("Name: %s \nUrl: %s \nColor: %s\n", job.Name, job.Url, job.Color)
		}
	},
}

func init() {
	rootCmd.AddCommand(listJobCmd)
	// listJobCmd.Flags().StringVarP(&jenkinsURL, "url", "u", "http://localhost:8080", "Jenkins URL")
	// listJobCmd.Flags().StringVarP(&username, "username", "n", "", "Jenkins username")
	// listJobCmd.Flags().StringVarP(&apiToken, "apiToken", "t", "111054ed5d1aff1bc2757e64f081975d08", "Jenkins token")
	// _ = listJobCmd.MarkFlagRequired("username")
	// _ = listJobCmd.MarkFlagRequired("apiToken")
}
