package cmd

import (
	"fmt"

	"github.com/kings/ctfmini/internal/hashing"
	"github.com/spf13/cobra"
)

var hashAlgo string
var hashIn string

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Hash utilities",
}

var hashSumCmd = &cobra.Command{
	Use:   "sum",
	Short: "Print digest of file",
	RunE: func(cmd *cobra.Command, args []string) error {
		sum, err := hashing.DigestFile(hashIn, hashAlgo)
		if err != nil {
			return err
		}
		fmt.Println(sum)
		return nil
	},
}

func init() {
	hashCmd.AddCommand(hashSumCmd)
	hashSumCmd.Flags().StringVar(&hashAlgo, "algo", "sha256", "md5|sha1|sha256|sha512")
	hashSumCmd.Flags().StringVar(&hashIn, "in", "", "input file (required)")
	_ = hashSumCmd.MarkFlagRequired("in")
}
