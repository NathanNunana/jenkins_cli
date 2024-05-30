package cmd

import (
	"fmt"
	"log"

	"github.com/NathanNunana/jenkins_cli/internal/jenkins"
	"github.com/spf13/cobra"
)

var jobName string

var createJobCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a jenkins job",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := jenkins.CreateJob(jobName)
		if err != nil {
			log.Fatalf("failed to create job, %v", err)
		}
		fmt.Printf("created job, %s", jobName)
	},
}

func init() {
	rootCmd.AddCommand(createJobCmd)
	createJobCmd.Flags().StringVarP(&jobName, "name", "n", "", "Jenkins Job Name")
	_ = createJobCmd.MarkFlagRequired("name")
}
