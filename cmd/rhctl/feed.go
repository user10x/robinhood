package main

import (
	"context"
	"encoding/json"
	"fmt"
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
	listFeedCmd.Flags().String("output", "json", "json/table")

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
		category := ""
		if len(args) >= 0 {
			category = "popular_lists"
			fmt.Println(category)
		}

		f, err := feed.ListFeeds(ctx, c.DoraHost, fmt.Sprintf("/feed/?%s", category), ar.AccessToken)
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
		e := json.NewEncoder(output)
		_ = e.Encode(feeds)
	default:
		table := TableWriter(output)

		table.SetHeader([]string{"id", "category", "displayLabel", "templates", "contentType", "relatedNames"})
		for _, f := range feeds {
			contentType := ""
			relatedNames := ""
			for _, content := range f.Contents {
				contentType = content.ContentType
				//fmt.Println(content)
				fmt.Println(content.Data.DisplayName)
				for _, related := range content.Data.RelatedInstruments {
					//strings.Join(related.Name, " ")
					fmt.Println(related.Name)
					relatedNames = related.Name
					//strings.Join(related.Name, " ")

				}

			}
			table.Append([]string{
				f.ID,
				f.Category,
				f.DisplayLabel,
				strings.Join(f.Templates, " "),
				contentType,
				relatedNames,
			})
		}
		table.Render()
	}
}
