package cmd

import (
	"registryhub/source"

	"github.com/spf13/cobra"
)


var initCommand = &cobra.Command {
	Use: "init",
	Run: func(cmd *cobra.Command, args []string) {
    source.Run()
  },
}

func init() {
	rootCmd.AddCommand(initCommand)
}