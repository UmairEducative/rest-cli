package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update user data",
	Long:  `Update the user data.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
		user := User{Username: username, Password: password}
		fmt.Println(user)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
