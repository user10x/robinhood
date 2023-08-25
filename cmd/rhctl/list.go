package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"robinhood/pkg/instruments"
)

var instrumentsCmd = &cobra.Command{
	Use:   "instruments",
	Short: "instruments management commands",
}

func init() {
	instrumentsCmd.AddCommand(listInstrumentsCmd)
	listInstrumentsCmd.Flags().String("order-by", "createdAtTime", "order by column")
}

var listInstrumentsCmd = &cobra.Command{
	Use:   "list",
	Short: "list instruments",
	RunE: func(cmd *cobra.Command, args []string) error {

		// process this limit client side because we need to filter deltas out
		limit, err := cmd.Flags().GetInt("limit")
		if err != nil {
			return err
		}
		fmt.Println(limit)
		includeDeltas, err := cmd.Flags().GetBool("include-deltas")
		if err != nil {
			return err
		}
		fmt.Println(includeDeltas)

		orderBy, err := cmd.Flags().GetString("order-by")
		if err != nil {
			return err
		}
		fmt.Println(orderBy)

		releaseType, _ := cmd.Flags().GetString("type")
		fmt.Println(releaseType)
		buildType, _ := cmd.Flags().GetString("build-type")
		fmt.Println(buildType)

		ctx := context.Background()

		ar := GetToken(ctx)

		_, err = instruments.ListInstruments(ctx, c.Host, ar.AccessToken)

		return nil
	}}
