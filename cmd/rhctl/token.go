package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"robinhood/pkg/auth"
)

var tokenCmd = &cobra.Command{
	Use:   "login",
	Short: "Send login request",
	Run: func(cmd *cobra.Command, args []string) {
		// Create a JSON object

		ctx := context.Background()

		ar := GetToken(ctx)

		fmt.Println(ar)

	}}

func GetToken(ctx context.Context) *auth.AuthResponse {

	auth := &auth.AuthRequest{
		Username:    c.Username,
		Password:    c.Password,
		GrantType:   c.GrantType,
		ClientId:    c.ClientId,
		ExpiresIn:   86400,
		DeviceToken: c.DeviceToken,
		Scope:       c.Scope,
		//MfaCode:     "",
	}

	authResponse, err := auth.GetToken(ctx, c.Host)

	// save auth info to a json file
	if err != nil {
		fmt.Println("error getting token", err)
		panic(err)
	}
	return authResponse

}
