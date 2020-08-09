package cmd

import (
	"log"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	count   uint
	workers uint
	out     string
)

func init() {
	rootCmd.PersistentFlags().UintVarP(&count, "count", "n", 1, "elements to generate")
	rootCmd.PersistentFlags().UintVarP(&workers, "workers", "w", uint(runtime.NumCPU()), "generation concurrency")
	rootCmd.PersistentFlags().StringVarP(&out, "output", "o", "", "output file")
}

var rootCmd = &cobra.Command{
	Use:   "nohi",
	Short: "NOHI - Not Only Human-friendly Identity generator CLI tool and lib",
	Long:  "Fast command-line tool and library for generation human ids, uuids and passwords.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
