package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
	"os"
	"robinhood/pkg/instruments"
	"robinhood/pkg/models/client/robinhood"
	"strconv"
	"time"
)

var instrumentsCmd = &cobra.Command{
	Use:   "instruments",
	Short: "instruments management commands",
}

func init() {
	instrumentsCmd.AddCommand(listInstrumentsCmd)
	instrumentsCmd.AddCommand(listInstrumentsByListType)
	listInstrumentsCmd.Flags().Int("limit", 10, "limit output")
	listInstrumentsCmd.Flags().String("order-by", "desc", "order by column")
	listInstrumentsByListType.Flags().String("type", "t", "type of the list")
	listInstrumentsByListType.Flags().StringP("output", "o", "", "output format json")

	instrumentsCmd.AddCommand(getRatingsCmd)
	getRatingsCmd.Flags().String("id", "", "get rating by id")

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
		printInstruments(cmd, instruments, limit)

		return nil
	}}

func printInstruments(cmd *cobra.Command, instruments *robinhood.Instruments, limit int) {
	outputFormat, _ := cmd.Flags().GetString("output")
	output := os.Stdout
	switch outputFormat {
	case "json":
		e := json.NewEncoder(output)
		_ = e.Encode(instruments)
	default:
		table := TableWriter(output)
		table.SetHeader([]string{"id", "name", "symbol", "country", "listDate", "market", "tradeRatio", "test", "spac", "ipoasupportdsp"})
		for _, inst := range instruments.Results[:limit] {
			table.Append([]string{
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

var listToIdMap = map[string]string{
	"dailyMovers": "eddbebe5-34cc-4df1-953c-d3e3cb55bc19",
}

var listInstrumentsByListType = &cobra.Command{
	Use:   "listSpecific",
	Short: "list instruments gives specific list",
	RunE: func(cmd *cobra.Command, args []string) error {

		listType, err := cmd.Flags().GetString("type")
		if err != nil {
			return err
		}
		_, ok := listToIdMap[listType]

		if !ok {
			fmt.Println(ok)
		}

		ctx := context.Background()

		ar, err := getApiToken(ctx)

		if err != nil {
			return err
		}

		list, err := instruments.ListInstrumentsByListId(ctx, c.Host, ar.AccessToken)
		if err != nil {
			return err
		}
		//printInstruments(cmd, instruments.Results[:limit])
		printListByType(cmd, list)
		return nil
	}}

func printListByType(cmd *cobra.Command, result *robinhood.ListResultsByType) {
	outputFormat, _ := cmd.Flags().GetString("output")
	output := os.Stdout
	switch outputFormat {
	case "json":
		e := json.NewEncoder(output)
		_ = e.Encode(result)
	default:
		table := TableWriter(output)
		table.SetHeader([]string{"id", "name", "symbol", "state", "objId", "type", "weight", "holdings", "cap", "%", "OpenPos", "updated", "created", "usa", "uk"})
		for _, r := range result.Results {
			table.Append([]string{
				r.Id,
				r.Name,
				r.Symbol,
				r.State,
				r.ObjectId,
				r.ObjectType,
				r.Weight,
				strconv.FormatBool(r.Holdings),
				fmt.Sprintf("%f", r.MarketCap),
				fmt.Sprintf("%f", r.OneDayDollarChange),
				fmt.Sprintf("%d", r.OpenPositions),
				//fmt.Sprintf("%s.", humanize.Time(r.UpdatedAt)),
				//fmt.Sprintf("%s.", humanize.Time(r.CreatedAt)),
				humanize.RelTime(r.UpdatedAt, time.Now(), "earlier", "later"),
				humanize.RelTime(r.CreatedAt, time.Now(), "earlier", "later"),
				r.OwnerType,
				r.UkTradability,
				r.UsTradability,

				//strconv.FormatBool(r.IsActive),
			})
		}
		table.Render()
	}
}

var getRatingsCmd = &cobra.Command{
	Use:   "get",
	Short: "get instruments by id/symbol",
	RunE: func(cmd *cobra.Command, args []string) error {

		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		if id == "" {
			return nil
		}

		ctx := context.Background()

		ar, err := getApiToken(ctx)

		if err != nil {
			return err
		}

		rating, err := instruments.GetRatingById(ctx, c.Host, ar.AccessToken, id)

		fmt.Println(rating)

		if err != nil {
			return err
		}
		return nil
	}}
