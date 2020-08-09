package cmd

import (
	"github.com/spf13/cobra"

	"gitlab.com/unitto/nohi/internal/utils"
	"gitlab.com/unitto/nohi/pkg/adapter"
	"gitlab.com/unitto/nohi/pkg/hid"
)

var (
	sep         string
	suffix      uint
	suffixDigit bool
)

func init() {
	rootCmd.AddCommand(hidCmd)
	hidCmd.Flags().StringVarP(&sep, "separator", "s", "_", "separator characters")
	hidCmd.Flags().UintVarP(&suffix, "length", "l", 5, "suffix length")
	hidCmd.Flags().BoolVarP(&suffixDigit, "digit", "d", true, "enable digits in suffix")
}

var hidCmd = &cobra.Command{
	Use:   "hid",
	Short: "Generate human-friendly ID",
	Run: func(cmd *cobra.Command, args []string) {
		g := hid.New()
		g.SetSeparator(sep)
		g.SetLength(suffix)
		g.EnableDigit(suffixDigit)
		g.EnableUpper(false)
		g.EnableSpecial(false)

		a := adapter.New(g)
		a.SetWorkerCount(workers)

		res := make(chan string, count)
		go a.Bulk(count, res)
		utils.Output(res, out)
	},
}
