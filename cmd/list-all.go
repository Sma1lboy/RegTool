package cmd

import (
	"registryhub/console"
	"registryhub/source"

	"github.com/spf13/cobra"
)

var listAllCommand = &cobra.Command{
	Use:   "list-all",
	Short: "List All remote sources",

	Run: func(cmd *cobra.Command, args []string) {

		sources, err := source.GetRemoteRegistrySources()
		if err != nil {
			console.Error("Failed to get remote sources:", err.Error())
			return
		}
		source.PrintSources(sources)
	},
}

func init() {
	rootCmd.AddCommand(listAllCommand)
}
