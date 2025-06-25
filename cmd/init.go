package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/ignis-sh/goignis/pkg"
	"github.com/spf13/cobra"
)

var (
	flagInitDaemon bool
	flagInitConfig string
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:               "init",
	Short:             "Initialize Ignis",
	Args:              cobra.NoArgs,
	ValidArgsFunction: SimpleCompletionFunc(nil),
	Run: func(cmd *cobra.Command, args []string) {
		pid, err := pkg.InitIgnis(context.Background(), flagInitConfig, flagInitDaemon)
		if err != nil {
			cmd.PrintErrln("Failed:", err)
			os.Exit(1)
			return
		}
		if flagInitDaemon && !RenderJSONIfFlagged(pid) {
			fmt.Println("Running ignis as a daemon, pid:", pid)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().BoolVarP(&flagInitDaemon, "daemon", "d", false, "Run as a daemon")
	initCmd.Flags().StringVarP(&flagInitConfig, "config", "c", "", "Path to the configuration file (default: $HOME/.config/ignis/config.py)")
}
