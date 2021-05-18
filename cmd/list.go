package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available users",
	Long:  `The list command lists all available users.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		endpoint := "/getall"

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

		c := &http.Client{
			Timeout: 15 * time.Second,
		}

		resp, err := c.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Println(resp)
			return
		}
		defer resp.Body.Close()

		var users = []User{}
		SliceFromJSON(users, resp.Body)
		data, err := PrettyJSON(users)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Data: ", data)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
