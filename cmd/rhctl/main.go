package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"robinhood/config"
)

var c *config.Config

func init() {
	c = config.NewConfig()
	rootCmd.AddCommand(tokenCmd)
	rootCmd.AddCommand(instrumentsCmd)

}

var rootCmd = &cobra.Command{
	Use:           "rhctl <command> [command-flags] [command-args]",
	Short:         "rhctl is a simple client interface into the robhinhood system",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
