package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new user",
	Long:  `Add a new user to the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		user := User{Username: username, Password: password}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
