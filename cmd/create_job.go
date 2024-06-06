package cmd

import (
	"fmt"
	"log"

	"github.com/NathanNunana/jenkins_cli/internal/jenkins"
	"github.com/spf13/cobra"
)

var (
	jobName         string
	repoURL         string
	branchName      string
	jenkinsfilePath string
	credentialsId   string
	jobType         string
)

var createJobCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a jenkins job",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		created_job, err := jenkins.CreateJob(jobName, repoURL, jenkinsfilePath, branchName, credentialsId, jobType)
		if err != nil {
			log.Fatalf("failed to create job, %v", err)
		}

		fmt.Printf("created job, %s", jobName)
		fmt.Println(created_job.Name)
	},
}

func init() {
	rootCmd.AddCommand(createJobCmd)
	createJobCmd.Flags().StringVarP(&jobName, "name", "n", "", "Jenkins job name")
	createJobCmd.Flags().StringVarP(&repoURL, "repo", "r", "", "Repository URL")
	createJobCmd.Flags().StringVarP(&credentialsId, "credentials", "c", "", "credential to create job")
	// set defaults for jenkinsfile path,branch name and job type
	createJobCmd.PersistentFlags().StringVarP(&branchName, "branch", "b", "main", "Branch Name")
	createJobCmd.PersistentFlags().StringVarP(&jenkinsfilePath, "file", "f", "Jenkinsfile", "File path of Jenkins script")
	createJobCmd.PersistentFlags().StringVarP(&jobType, "type", "t", "multi-branch", "Specify Jenkins job type (multi-branch,freestyle)")

	_ = createJobCmd.MarkFlagRequired("name")
	_ = createJobCmd.MarkFlagRequired("repo")
	_ = createJobCmd.MarkFlagRequired("credentials")
	_ = createJobCmd.MarkFlagRequired("jenkinsfilePath")
	_ = createJobCmd.MarkFlagRequired("jobType")
}
