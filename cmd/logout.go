package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout for user",
	Long:  `Logging out a user from the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logout called")
		user := User{Username: username, Password: password}
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
