package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/ignis-sh/goignis/pkg"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:               "init",
	Short:             "Initialize Ignis",
	Args:              cobra.NoArgs,
	ValidArgsFunction: argsCompletionFunc(nil),
	Run: func(cmd *cobra.Command, args []string) {
		config, _ := cmd.Flags().GetString("config")
		daemon, _ := cmd.Flags().GetBool("daemon")
		pid, err := pkg.InitIgnis(context.Background(), config, daemon)
		if err != nil {
			cmd.PrintErrln("Failed:", err)
			os.Exit(1)
			return
		}
		if daemon && !renderResultIfJSON(pid) {
			fmt.Println("Running ignis as a daemon, pid:", pid)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().BoolP("daemon", "d", false, "Run as a daemon")
	initCmd.Flags().StringP("config", "c", "", "Path to the configuration file (default: $HOME/.config/ignis/config.py)")
}
