package cmd

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search the database",
	Long: `Search the database for a user, identified by a user name.
	The command returns the full record of the user.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
		endpoint := "/getid"
		user := User{Username: username, Password: password}
		fmt.Println(user)

		// bytes.Buffer is both a Reader and a Writer
		buf := new(bytes.Buffer)
		err := user.ToJSON(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		req, err := http.NewRequest("GET", SERVER+PORT+endpoint, buf)
		if err != nil {
			fmt.Println("GetAll – Error in req: ", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
