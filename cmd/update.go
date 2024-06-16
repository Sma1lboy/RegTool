// update applications' registry to somewhere
package cmd

import (
	"fmt"
	"os"
	"registryhub/source"

	"github.com/spf13/cobra"
)

var region string
var app string

var updateCmd = &cobra.Command{
	Use:   "update [region]",
	Short: "Update sources to the specified region registry",
	Long: `Update sources managed by RegistryHub to use the registry of the specified region.
   If --app is provided, update only the specified app. Valid regions are 'cn' for China and 'us' for the United States.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		region := args[0]
		var err error
		if app != "" {
			// err = updateSoftwareRegistry(app, region)
		} else {
			if !source.ChangeAllRegistry(region) {
				err = fmt.Errorf("failed to update all sources to the %s registry", region)
			}

		}

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		if app != "" {
			fmt.Printf("Successfully updated %s to the %s registry.\n", app, region)
		} else {
			fmt.Printf("Successfully updated all sources to the %s registry.\n", region)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVar(&app, "app", "", "Specific app to update registry for (optional)")
}
