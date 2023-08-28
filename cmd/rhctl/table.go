package main

import (
	"github.com/olekukonko/tablewriter"
	"io"
)

func TableWriter(out io.Writer) *tablewriter.Table {
	table := tablewriter.NewWriter(out)
	//table.SetBorder(true)
	//table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	//table.SetHeaderLine(true)
	//table.SetColumnSeparator("|")
	//table.SetRowLine(true)
	//table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoWrapText(true)
	//table.NumLines()
	//table.SetRowLine(true)
	return table
}
