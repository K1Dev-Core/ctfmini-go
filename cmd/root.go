package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ctfmini",
	Short: "CTFmini - tiny CTF toolkit CLI (Go single binary)",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// namespaces
	rootCmd.AddCommand(encCmd)
	rootCmd.AddCommand(classicCmd)
	rootCmd.AddCommand(hashCmd)
	rootCmd.AddCommand(utilCmd)
}
