package cmd

import (
	"github.com/spf13/cobra"

	"gitlab.com/unitto/nohi/internal/utils"
	"gitlab.com/unitto/nohi/pkg/adapter"
	"gitlab.com/unitto/nohi/pkg/uuid4"
)

func init() {
	rootCmd.AddCommand(uuid4Cmd)
}

var uuid4Cmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate UUID (version 4)",
	Run: func(cmd *cobra.Command, args []string) {
		a := adapter.New(uuid4.New())
		a.SetWorkerCount(workers)

		res := make(chan string, count)
		go a.Bulk(count, res)
		utils.Output(res, out)
	},
}
