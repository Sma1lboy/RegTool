// update applications' registry to somewhere
package cmd

import (
	"fmt"
	"os"
	"registryhub/source"

	"github.com/spf13/cobra"
)

var region string
var target string

var updateCmd = &cobra.Command{
	Use:   "update [region]",
	Short: "Update sources to the specified region registry",
	Long: `Update sources managed by RegistryHub to use the registry of the specified region.
   If --target is provided, update only the specified app. Valid regions are 'cn' for China and 'us' for the United States.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		region := args[0]
		var err error
		if target != "" {
			err2 := source.UpdateRegistry(target, region)
			if err2 != nil {
				err = fmt.Errorf("invalid app name to update %s to the %s registry", target, region)
			}

		} else {
			if !source.ChangeAllRegistry(region) {
				err = fmt.Errorf("failed to update all sources to the %s registry", region)
			}

		}

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		if target != "" {
			fmt.Printf("Successfully updated %s to the %s registry.\n", target, region)
		} else {
			fmt.Printf("Successfully updated all sources\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&target, "target", "t", "", "Specific target to update registry for (optional)")
}
