package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"robinhood/pkg/auth"
)

var tokenCmd = &cobra.Command{
	Use:   "myapp",
	Short: "Send an HTTP request with JSON body",
	Run: func(cmd *cobra.Command, args []string) {
		// Create a JSON object

		ctx := context.Background()

		auth := &auth.AuthRequest{
			Username:    "nipun.chawla786@gmail.com",
			Password:    "Dallas007$$",
			GrantType:   "password",
			ClientId:    "c82SH0WZOsabOXGP2sxqcj34FxkvfnWRZBKlBjFS",
			ExpiresIn:   86400,
			DeviceToken: "c547515b-67a5-4e60-9768-8e211915a84c",
			Scope:       "internal",
			//MfaCode:     "",
		}

		_, err := auth.GetToken(ctx)

		// save auth info to a json file
		if err != nil {
			panic(err)
		}

	}}

func init() {
	//var name string
	//var age int
	//tokenCmd.Flags().StringVar(&name, "name", "", "Name of person")
	//tokenCmd.Flags().IntVar(&age, "name", 0, "Age of person")
}

func main() {
	if err := tokenCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

// Create a new request with JSON body
