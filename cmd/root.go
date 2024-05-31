// cmd/root.go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "registryhub",
	Short: "RegistryHub is a very fast registry manager",
	Long: `RegistryHub is a highly efficient tool designed for managing software
   download sources with ease and speed. Built with love by spf13 and friends
   using Go, RegistryHub simplifies the process of handling and organizing software repositories.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run hugo...")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
