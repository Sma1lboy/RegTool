package cmd

import (
	"registryhub/console"
	"registryhub/source"

	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List current sources using",

	Run: func(cmd *cobra.Command, args []string) {

		// m := source.ReadBackup()
		m, err := source.GetLocalSourcesMap()
		if err != nil {
			console.Error("Failed to get local sources:", err.Error())
			return
		}

		source.PrintSources(m)

	},
}

func init() {
	rootCmd.AddCommand(listCommand)
}
