package cmd

import (
	"fmt"
	"registryhub/source"

	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use: "init",
	Run: func(cmd *cobra.Command, args []string) {
		sources := source.Run()
		for v, s := range sources {
			fmt.Println(v, s)
		}

	},
}

func init() {
	rootCmd.AddCommand(initCommand)
}
