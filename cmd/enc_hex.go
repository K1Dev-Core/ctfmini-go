package cmd

import (
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
)

var encHexDec bool

var encHexCmd = &cobra.Command{
	Use:   "hex <text-or-hex>",
	Short: "Hex encode/decode",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		in := args[0]
		if encHexDec {
			b, err := hex.DecodeString(in)
			if err != nil {
				return err
			}
			fmt.Print(string(b))
			return nil
		}
		fmt.Print(hex.EncodeToString([]byte(in)))
		return nil
	},
}

func init() {
	encCmd.AddCommand(encHexCmd)
	encHexCmd.Flags().BoolVarP(&encHexDec, "dec", "d", false, "decode")
}
