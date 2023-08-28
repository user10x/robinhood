package main

import (
	"context"
	"encoding/json"
	"github.com/spf13/cobra"
	"os"
	"robinhood/pkg/models/client/robinhood"
	"robinhood/pkg/orders"
	"strconv"
)

var orderCmd = &cobra.Command{
	Use:   "orders",
	Short: "orders management commands",
}

func init() {
	orderCmd.AddCommand(listOrdersCmd)
}

var listOrdersCmd = &cobra.Command{
	Use:   "list",
	Short: "list orders",
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()

		ar, err := getApiToken(ctx)

		if err != nil {
			return err
		}

		o, err := orders.ListOrders(ctx, c.Host, "orders/", ar.AccessToken)

		if err != nil {
			return err
		}
		printOrders(cmd, o)

		return nil
	}}

func printOrders(cmd *cobra.Command, orders []robinhood.Order) {
	outputFormat, _ := cmd.Flags().GetString("output")
	output := os.Stdout
	switch outputFormat {
	case "json":
		e := json.NewEncoder(output)
		_ = e.Encode(orders)
	default:
		table := TableWriter(output)
		table.SetHeader([]string{"id", "inst", "fees", "cumulativeQnt", "timeInforce", "qnt", "state", "trigger", "created", "amt", "ipoAccessOrder", "ipoPriceFinalized"})
		for _, o := range orders {
			amt := ""
			if o.DollarBasedAmount != nil {
				amt = o.DollarBasedAmount.Amount
			}
			table.Append([]string{
				o.Id,
				o.Instrument,
				//o.Url,
				o.Fees,
				o.CumulativeQuantity,
				o.TimeInForce,
				o.Quantity,
				o.State,
				o.Trigger,
				o.CreatedAt.Format("January 2, 2006"),
				amt,
				strconv.FormatBool(o.IsIpoAccessOrder),
				strconv.FormatBool(o.IsIpoAccessPriceFinalized),
				//string(*o.AveragePrice),

				//r.InternalHaltReason,
				//inst.DayTradeRatio,
				//strconv.FormatBool(inst.IsTest),
				//strconv.FormatBool(inst.IsSpac),
				//strconv.FormatBool(inst.IpoAccessSupportsDsp),
				//r.CreatedAtTime.AsTime().Format(time.RFC3339),
				//strconv.FormatBool(r.IsActive),
			})
		}
		table.Render()
	}
}
