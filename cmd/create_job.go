package cmd

import (
	"github.com/spf13/cobra"
)

var jobName string

var createJobCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a jenkins job",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// url := fmt.Sprintf("%s/createItem/", jenkinsURL)
	},
}

func init() {
	rootCmd.AddCommand(createJobCmd)
}
