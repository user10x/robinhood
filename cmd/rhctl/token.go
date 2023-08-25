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
		ctx := context.Background()

		_, err := GetToken(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Printf("sucessfully logged in!\n")
	}}

func GetToken(ctx context.Context) (*auth.AuthResponse, error) {

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
	if err != nil {
		return nil, err
	}
	return authResponse, nil

}
