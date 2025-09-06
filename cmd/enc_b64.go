package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

var encB64Dec bool

var encB64Cmd = &cobra.Command{
	Use:   "b64 <text>",
	Short: "Base64 encode/decode",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		in := args[0]
		if encB64Dec {
			out, err := base64.StdEncoding.DecodeString(in)
			if err != nil {
				return err
			}
			fmt.Print(string(out))
			return nil
		}
		fmt.Print(base64.StdEncoding.EncodeToString([]byte(in)))
		return nil
	},
}

var encCmd = &cobra.Command{
	Use:   "enc",
	Short: "Encoders/decoders (b64, hex, url)",
}

func init() {
	encCmd.AddCommand(encB64Cmd)
	encB64Cmd.Flags().BoolVarP(&encB64Dec, "dec", "d", false, "decode")
}
