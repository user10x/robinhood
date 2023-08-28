package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"robinhood/pkg/instruments"
	"robinhood/pkg/models/client/robinhood"
	"strconv"
)

var instrumentsCmd = &cobra.Command{
	Use:   "instruments",
	Short: "instruments management commands",
}

func init() {
	instrumentsCmd.AddCommand(listInstrumentsCmd)
	listInstrumentsCmd.Flags().Int("limit", 10, "limit output")
	listInstrumentsCmd.Flags().String("order-by", "desc", "order by column")
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

		orderBy, err := cmd.Flags().GetString("order-by")
		if err != nil {
			return err
		}
		fmt.Println(orderBy)

		ctx := context.Background()

		ar, err := getApiToken(ctx)

		if err != nil {
			return err
		}

		instruments, err := instruments.ListInstruments(ctx, c.Host, ar.AccessToken)

		if err != nil {
			return err
		}
		if len(instruments.Results) < limit {
			limit = len(instruments.Results)
		}
		printInstruments(cmd, instruments.Results[:limit])

		return nil
	}}

func printInstruments(cmd *cobra.Command, instruments []robinhood.Instrument) {
	outputFormat, _ := cmd.Flags().GetString("output")
	output := os.Stdout
	switch outputFormat {
	case "json":
		e := newJSONEncoder(output)
		_ = e.Encode(instruments)
	default:
		table := TableWriter(output)
		table.SetHeader([]string{"id", "name", "symbol", "country", "listDate", "market", "tradeRatio", "test", "spac", "ipoasupportdsp"})
		for _, inst := range instruments {
			table.Append([]string{
				//r.Results,
				inst.Id,
				inst.Name,
				inst.Symbol,
				inst.Country,
				inst.ListDate,
				inst.Market,
				//r.InternalHaltReason,
				inst.DayTradeRatio,
				strconv.FormatBool(inst.IsTest),
				strconv.FormatBool(inst.IsSpac),
				strconv.FormatBool(inst.IpoAccessSupportsDsp),
				//r.CreatedAtTime.AsTime().Format(time.RFC3339),
				//strconv.FormatBool(r.IsActive),
			})
		}
		table.Render()
	}
}
