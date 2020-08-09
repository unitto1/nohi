package cmd

import (
	"github.com/spf13/cobra"

	"gitlab.com/unitto/nohi/internal/utils"
	"gitlab.com/unitto/nohi/pkg/adapter"
	"gitlab.com/unitto/nohi/pkg/pwd"
)

var (
	length  uint
	upper   bool
	digit   bool
	special bool
)

func init() {
	rootCmd.AddCommand(pwdCmd)
	pwdCmd.Flags().UintVarP(&length, "length", "l", 32, "password length")
	pwdCmd.Flags().BoolVarP(&upper, "upper", "u", false, "enable upper characters")
	pwdCmd.Flags().BoolVarP(&digit, "digit", "d", true, "enable digit characters")
	pwdCmd.Flags().BoolVarP(&special, "special", "s", false, "enable special characters")
}

var pwdCmd = &cobra.Command{
	Use:   "pwd",
	Short: "Generate password",
	Run: func(cmd *cobra.Command, args []string) {
		g := pwd.New()
		g.SetLength(length)
		g.EnableUpper(upper)
		g.EnableDigit(digit)
		g.EnableSpecial(special)

		a := adapter.New(g)
		a.SetWorkerCount(workers)

		res := make(chan string, count)
		go a.Bulk(count, res)
		utils.Output(res, out)
	},
}
