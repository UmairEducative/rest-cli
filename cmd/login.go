package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login into the system",
	Long:  `The login command logins an existing user into the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
		user := User{Username: username, Password: password}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
