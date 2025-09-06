package cmd

import (
	"fmt"

	"github.com/kings/ctfmini/internal/classical"
	"github.com/spf13/cobra"
)

var rotN int

var rotCmd = &cobra.Command{
	Use:   "rot <text>",
	Short: "ROT-N (default 13)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(classical.RotN(args[0], rotN))
	},
}

var classicCmd = &cobra.Command{
	Use:   "classic",
	Short: "Classical ciphers (rot, xor, vigenere)",
}

func init() {
	classicCmd.AddCommand(rotCmd)
	rotCmd.Flags().IntVarP(&rotN, "n", "n", 13, "shift amount")
}
