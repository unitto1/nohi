package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var Version = "undefined"
var Commit = "undefined"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of NOHI",
	Long:  "All software has a versions. This is NOHI one.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version:\t%s\n", Version)
		fmt.Printf("Commit: \t%s\n", Commit)
	},
}
