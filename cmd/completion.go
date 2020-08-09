package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completionCmd)
}

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish]",
	Short: "Generate auto-completion script",
	Long: `To load auto-completion:

Bash:
Linux:
$ nohi completion bash > /etc/bash_completion.d/nohi
MacOS:
$ nohi completion bash > /usr/local/etc/bash_completion.d/nohi

Zsh:
$ nohi completion zsh > "${fpath[1]}/nohi"

Fish:
$ nohi completion fish > ~/.config/fish/completions/nohi.fish
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		switch args[0] {
		case "bash":
			err = cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			err = cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			err = cmd.Root().GenFishCompletion(os.Stdout, true)
		}

		if err != nil {
			log.Fatal(err)
		}
	},
}
