package main

import (
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"io"
)

// newJSONEncoder creates a custom json encoder that makes the json output a bit more
// readable
func newJSONEncoder(w io.Writer) *json.Encoder {
	return newJSONEncoder(w)
}

func TableWriter(out io.Writer) *tablewriter.Table {
	table := tablewriter.NewWriter(out)
	table.SetBorder(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(true)
	table.SetColumnSeparator("|")
	table.SetRowLine(true)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoWrapText(true)
	table.SetRowLine(true)
	return table
}
