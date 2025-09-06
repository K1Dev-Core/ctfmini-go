package cmd

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/kings/ctfmini/internal/util"
	"github.com/spf13/cobra"
)

var xorKey string
var xorHexKey bool
var xorIn string
var xorOut string

var xorCmd = &cobra.Command{
	Use:   "xor [text]",
	Short: "XOR bytes with key (supports --in/--out, --hexkey)",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var data []byte
		var err error
		if xorIn != "" {
			data, err = os.ReadFile(xorIn)
			if err != nil {
				return err
			}
		} else if len(args) == 1 {
			data = []byte(args[0])
		} else {
			// read stdin
			data, err = util.ReadAllStdin()
			if err != nil {
				return err
			}
		}

		var key []byte
		if xorHexKey {
			key, err = hex.DecodeString(xorKey)
			if err != nil {
				return err
			}
		} else {
			key = []byte(xorKey)
		}
		out := util.XOR(data, key)

		if xorOut != "" {
			return os.WriteFile(xorOut, out, 0o644)
		}
		_, _ = os.Stdout.Write(out)
		return nil
	},
}

func init() {
	classicCmd.AddCommand(xorCmd)
	xorCmd.Flags().StringVarP(&xorKey, "key", "k", "", "key (required)")
	_ = xorCmd.MarkFlagRequired("key")
	xorCmd.Flags().BoolVar(&xorHexKey, "hexkey", false, "key is hex")
	xorCmd.Flags().StringVar(&xorIn, "in", "", "input file")
	xorCmd.Flags().StringVar(&xorOut, "out", "", "output file")
}
