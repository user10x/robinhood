package main

import (
	"context"
	"github.com/spf13/cobra"
	"os"
	"robinhood/pkg/feed"
	"robinhood/pkg/models/client/robinhood"
	"strings"
)

var feedCmd = &cobra.Command{
	Use:   "feed",
	Short: "news feed management commands",
}

func init() {
	feedCmd.AddCommand(listFeedCmd)
}

var listFeedCmd = &cobra.Command{
	Use:   "list",
	Short: "list feeds",
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()

		ar, err := getApiToken(ctx)

		if err != nil {
			return err
		}

		f, err := feed.ListFeeds(ctx, c.DoraHost, ar.AccessToken)

		if err != nil {
			return err
		}
		printFeed(cmd, f)

		return nil
	}}

func printFeed(cmd *cobra.Command, feeds []robinhood.Feed) {
	outputFormat, _ := cmd.Flags().GetString("output")
	output := os.Stdout
	switch outputFormat {
	case "json":
		e := newJSONEncoder(output)
		_ = e.Encode(nil)
	default:
		table := TableWriter(output)

		table.SetHeader([]string{"id", ""})
		for _, f := range feeds {
			table.Append([]string{
				f.ID,
				f.DisplayLabel,
				f.URL,
				f.RankingVersion,
				strings.Join(f.Templates, " "),
			})
		}
		table.Render()
	}
}
