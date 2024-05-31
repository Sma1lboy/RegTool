package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of Registry Hub",
  Long:  `This is Registry Hub version`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Registry Hub v1.0 -- HEAD")
  },
}

func init() {
  rootCmd.AddCommand(versionCmd)
}
