package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search the database",
	Long: `Search the database for a user, identified by a user name.
	The command returns the full record of the user.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
		user := User{Username: username, Password: password}
		fmt.Println(user)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
