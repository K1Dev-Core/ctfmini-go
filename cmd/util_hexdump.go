package cmd

import (
	"fmt"
	"os"

	"github.com/kings/ctfmini/internal/util"
	"github.com/spf13/cobra"
)

var hexdIn string
var hexdWidth int

var hexdCmd = &cobra.Command{
	Use:   "hexdump",
	Short: "Simple hex + ASCII view",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := os.ReadFile(hexdIn)
		if err != nil {
			return err
		}
		fmt.Print(util.HexDump(data, hexdWidth))
		return nil
	},
}

var utilCmd = &cobra.Command{
	Use:   "util",
	Short: "Utilities (hexdump, strings)",
}

func init() {
	utilCmd.AddCommand(hexdCmd)
	hexdCmd.Flags().StringVar(&hexdIn, "in", "", "input file (required)")
	_ = hexdCmd.MarkFlagRequired("in")
	hexdCmd.Flags().IntVar(&hexdWidth, "width", 16, "bytes per line")
}
