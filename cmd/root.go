package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"user"`
	Password  string `json:"password"`
	LastLogin int64  `json:"lastlogin"`
	Admin     int    `json:"admin"`
	Active    int    `json:"active"`
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

	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("port"))
}
