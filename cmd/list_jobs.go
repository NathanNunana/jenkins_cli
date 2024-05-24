package cmd

import (
	"fmt"
	"log"

	"github.com/NathanNunana/jenkins_cli/internal/jenkins"
	"github.com/NathanNunana/jenkins_cli/internal/printer"
	"github.com/spf13/cobra"
)

var listJobCmd = &cobra.Command{
	Use:   "list",
	Short: "list all jenkins jobs",
	Long:  "Produces a list of all jenkins jobs created",
	Run: func(cmd *cobra.Command, args []string) {
		jobs, err := jenkins.GetJobs()
		if err != nil {
			log.Fatal(err)
		}
		// display the response
		jobPrinter := printer.PrintJob{Jobs: jobs}
		data, err := jobPrinter.List()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(data)
	},
}

func init() {
	rootCmd.AddCommand(listJobCmd)
}
