package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Get the time from the RESTful server",
	Long:  `This command mainly exists for making sure that the server works.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("time called")
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)
}
