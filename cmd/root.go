package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var SERVER string
var PORT string
var data string
var username string
var password string

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"user"`
	Password  string `json:"password"`
	LastLogin int64  `json:"lastlogin"`
	Admin     int    `json:"admin"`
	Active    int    `json:"active"`
}

func (p *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

var rootCmd = &cobra.Command{
	Use:   "rest-cli",
	Short: "A REST API client",
	Long:  `A Client for a RESTful server.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("username", "u", "username", "The username")
	rootCmd.PersistentFlags().StringP("password", "p", "1234", "The password")
	rootCmd.PersistentFlags().StringP("data", "d", "{}", "JSON Record")

	rootCmd.PersistentFlags().StringVar(&SERVER, "server", "http://localhost", "RESTful server hostname")
	rootCmd.PersistentFlags().StringVar(&PORT, "port", ":1234", "Port of RESTful Server")

	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("data", rootCmd.PersistentFlags().Lookup("data"))
	viper.BindPFlag("SERVER", rootCmd.PersistentFlags().Lookup("SERVER"))
	viper.BindPFlag("PORT", rootCmd.PersistentFlags().Lookup("PORT"))
}
