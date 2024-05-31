package cmd

import (
	"fmt"
	"registryhub/console"
	"registryhub/source"

	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List all sources",

	Run: func(cmd *cobra.Command, args []string) {
		// sources := source.ReadBackup()
		// for v, s := range sources {
		// 	console.Warning(v, s)
		// }
		sources, err := source.GetRemoteSources()
		if err != nil {
			console.Error("Failed to get remote sources:", err.Error())
			return
		}
		fmt.Println(sources)
		for region, registryRegion := range *sources {
			fmt.Println(region)
			for packageManager, urls := range registryRegion {
				console.Success(packageManager, urls[0])
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCommand)
}
