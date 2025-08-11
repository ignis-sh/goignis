package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ignis-sh/goignis/pkg"

	"github.com/spf13/cobra"
)

// runCommandCmd represents the run-command command
var runCommandCmd = &cobra.Command{
	Use:   "run-command",
	Short: "Run a custom command",
	Args:  cobra.MinimumNArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) (completions []cobra.Completion, directive cobra.ShellCompDirective) {
		if len(args) == 0 {
			// args[0] must be a command name
			directive = cobra.ShellCompDirectiveNoFileComp
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			// list available command names
			var err error
			completions, err = pkg.ListCommands(ctx)
			if err != nil {
				return
			}
		}
		return
	},
	Run: func(cmd *cobra.Command, args []string) {
		// run command `args[0]` with `args[1:]` as command args
		cmdError, output, err := pkg.RunCommand(context.Background(), args[0], args[1:])
		if err != nil {
			cmd.PrintErrln("Failed:", err)
			os.Exit(1)
			return
		}
		if cmdError != "" {
			cmd.PrintErrln(cmdError)
			os.Exit(1)
			return
		}
		if !RenderJSONIfFlagged(output) {
			if output != "" {
				fmt.Println(output)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCommandCmd)
}
