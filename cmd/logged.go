package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// loggedCmd represents the logged command
var loggedCmd = &cobra.Command{
	Use:   "logged",
	Short: "List add logged in users",
	Long:  `This command shows all logged in users.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logged called")
	},
}

func init() {
	rootCmd.AddCommand(loggedCmd)
}
