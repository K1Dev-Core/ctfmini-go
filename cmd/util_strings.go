package cmd

import (
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

var stringsIn string
var stringsMin int

var stringsCmd = &cobra.Command{
	Use:   "strings",
	Short: "Extract printable ASCII strings from a file",
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := os.ReadFile(stringsIn)
		if err != nil {
			return err
		}
		re := regexp.MustCompile(`[ -~]{` + fmt.Sprint(stringsMin) + `,}`)
		for _, m := range re.FindAll(b, -1) {
			fmt.Println(string(m))
		}
		return nil
	},
}

func init() {
	utilCmd.AddCommand(stringsCmd)
	stringsCmd.Flags().StringVar(&stringsIn, "in", "", "input file (required)")
	_ = stringsCmd.MarkFlagRequired("in")
	stringsCmd.Flags().IntVar(&stringsMin, "min", 4, "minimum length")
}
