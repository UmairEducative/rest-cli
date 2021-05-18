package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deleting users",
	Long:  `This command deletes existing users from the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
		user := User{Username: username, Password: password}
		fmt.Println(user)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
