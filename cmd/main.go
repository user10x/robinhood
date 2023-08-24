package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"robinhood/pkg/auth"
	"robinhood/pkg/instruments"
)

func getToken(ctx context.Context) *auth.AuthResponse {

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

	authResponse, err := auth.GetToken(ctx)

	// save auth info to a json file
	if err != nil {
		fmt.Println("error getting token", err)
		panic(err)
	}
	return authResponse

}

var tokenCmd = &cobra.Command{
	Use:   "login",
	Short: "Send login request",
	Run: func(cmd *cobra.Command, args []string) {
		// Create a JSON object

		ctx := context.Background()

		ar := getToken(ctx)

		fmt.Println(ar)

	}}

var listInstruments = &cobra.Command{
	Use:   "get all instruments",
	Short: "get all instruments",
	Run: func(cmd *cobra.Command, args []string) {
		// Create a JSON object

		ctx := context.Background()

		ar := getToken(ctx)

		const host = "https://api.robinhood.com/"

		_, err := instruments.ListInstruments(ctx, host, ar.AccessToken)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(instruments) pretty print json and table
	}}

func init() {
	//var name string
	//var age int
	//tokenCmd.Flags().StringVar(&name, "name", "", "Name of person")
	//tokenCmd.Flags().IntVar(&age, "name", 0, "Age of person")
}

func main() {
	//if err := tokenCmd.Execute(); err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

	if err := listInstruments.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

// Create a new request with JSON body
