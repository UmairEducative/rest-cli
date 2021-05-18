package cmd

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login into the system",
	Long:  `The login command logins an existing user into the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")

		endpoint := "/login"
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
	rootCmd.AddCommand(loginCmd)
}
