package cmd

import (
	"fmt"

	"github.com/kings/ctfmini/internal/classical"
	"github.com/spf13/cobra"
)

var vigKey string
var vigDec bool

var vigCmd = &cobra.Command{
	Use:   "vigenere <text>",
	Short: "Vigenère cipher enc/dec (letters only shifted)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(classical.Vigenere(args[0], vigKey, vigDec))
	},
}

func init() {
	classicCmd.AddCommand(vigCmd)
	vigCmd.Flags().StringVarP(&vigKey, "key", "k", "", "key (letters)")
	_ = vigCmd.MarkFlagRequired("key")
	vigCmd.Flags().BoolVarP(&vigDec, "dec", "d", false, "decrypt")
}
