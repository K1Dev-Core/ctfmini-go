package cmd

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
)

var encURLDec bool

var encURLCmd = &cobra.Command{
	Use:   "url <text>",
	Short: "URL encode/decode (application/x-www-form-urlencoded)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		in := args[0]
		if encURLDec {
			out, err := url.QueryUnescape(in)
			if err != nil {
				return err
			}
			fmt.Print(out)
			return nil
		}
		fmt.Print(url.QueryEscape(in))
		return nil
	},
}

func init() {
	encCmd.AddCommand(encURLCmd)
	encURLCmd.Flags().BoolVarP(&encURLDec, "dec", "d", false, "decode")
}
