package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// loggedCmd represents the logged command
var loggedCmd = &cobra.Command{
	Use:   "logged",
	Short: "List add logged in users",
	Long:  `This command shows all logged in users.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logged called")
		user := User{Username: username, Password: password}

		// bytes.Buffer is both a Reader and a Writer
		buf := new(bytes.Buffer)
		err := user.ToJSON(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		req, err := http.NewRequest("GET", SERVER+PORT+"/logged", buf)
		if err != nil {
			fmt.Println("LoggedUsers – Error in req: ", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		c := &http.Client{
			Timeout: 15 * time.Second,
		}

		resp, err := c.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Server response HTTP status code", resp.StatusCode)
			return
		}
		defer resp.Body.Close()

		var users = []User{}
		SliceFromJSON(&users, resp.Body)
		data, err := PrettyJSON(users)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("Data: ", data)
	},
}

func init() {
	rootCmd.AddCommand(loggedCmd)
}
