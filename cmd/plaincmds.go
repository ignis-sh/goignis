package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ignis-sh/goignis/pkg"

	"github.com/spf13/cobra"
)

func init() {
	renderWindowCmd := func(cmd *cobra.Command, found bool) {
		if !found {
			cmd.PrintErrln("No such window")
			os.Exit(1)
		}
	}

	// subcommands without args or return values
	AddPlainCmd("systeminfo", "Print system information", pkg.IgnisSystemInfo)
	AddPlainCmd("quit", "Quit Ignis", pkg.QuitIgnis)
	AddPlainCmd("reload", "Reload Ignis", pkg.ReloadIgnis)
	AddPlainCmd("inspector", "Open GTK Inspector", pkg.OpenInspector)
	// subcommands with no args but one single return value
	AddPlainCmdWithResult("list-windows", "List names of all windows", pkg.ListWindows, func(_ *cobra.Command, windows []string) {
		fmt.Println(strings.Join(windows, "\n"))
	})
	// subcommands with one single argument and one single return value
	AddPlainCmdWithArgResult("toggle-window", "Toggle a window", pkg.ToggleWindow, pkg.ListWindows, renderWindowCmd)
	AddPlainCmdWithArgResult("open-window", "Open a window", pkg.OpenWindow, pkg.ListWindows, renderWindowCmd)
	AddPlainCmdWithArgResult("close-window", "Close a window", pkg.CloseWindow, pkg.ListWindows, renderWindowCmd)
	// For more complicated subcommands, use `cobra-cli add xxx` instead
}

// AddPlainCmd adds a subcommand to root, without args or results.
func AddPlainCmd(use, short string, callback func(context.Context) error) {
	rootCmd.AddCommand(&cobra.Command{
		Use:               use,
		Short:             short,
		Args:              cobra.NoArgs,
		ValidArgsFunction: SimpleCompletionFunc(nil),
		Run: func(cmd *cobra.Command, args []string) {
			err := callback(context.Background())
			if err != nil {
				cmd.PrintErrln("Failed:", err)
				os.Exit(1)
				return
			}
		},
	})
}

// AddPlainCmdWithArg adds a subcommand to root, with no results but one single argument.
func AddPlainCmdWithArg(
	use, short string,
	callback func(context.Context, string) error,
	completion func(context.Context) ([]cobra.Completion, error),
) {
	rootCmd.AddCommand(&cobra.Command{
		Use:               use,
		Short:             short,
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: SimpleCompletionFunc(completion),
		Run: func(cmd *cobra.Command, args []string) {
			err := callback(context.Background(), args[0])
			if err != nil {
				cmd.PrintErrln("Failed:", err)
				os.Exit(1)
				return
			}
		},
	})
}

// AddPlainCmdWithResult adds a subcommand to root, with no args but one single result.
func AddPlainCmdWithResult[T any](
	use, short string,
	callback func(context.Context) (T, error),
	render func(*cobra.Command, T),
) {
	rootCmd.AddCommand(&cobra.Command{
		Use:               use,
		Short:             short,
		Args:              cobra.NoArgs,
		ValidArgsFunction: SimpleCompletionFunc(nil),
		Run: func(cmd *cobra.Command, args []string) {
			result, err := callback(context.Background())
			if err != nil {
				cmd.PrintErrln("Failed:", err)
				os.Exit(1)
				return
			}
			if !RenderJSONIfFlagged(result) {
				render(cmd, result)
			}
		},
	})
}

// AddPlainCmdWithArgResult adds a subcommand to root, with one argument and one result.
func AddPlainCmdWithArgResult[T any](
	use, short string,
	callback func(context.Context, string) (T, error),
	completion func(context.Context) ([]cobra.Completion, error),
	render func(*cobra.Command, T),
) {
	rootCmd.AddCommand(&cobra.Command{
		Use:               use,
		Short:             short,
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: SimpleCompletionFunc(completion),
		Run: func(cmd *cobra.Command, args []string) {
			result, err := callback(context.Background(), args[0])
			if err != nil {
				cmd.PrintErrln("Failed:", err)
				os.Exit(1)
				return
			}
			if !RenderJSONIfFlagged(result) {
				render(cmd, result)
			}
		},
	})
}

// SimpleCompletionFunc returns a shell completion func that provides fixed suggestions.
func SimpleCompletionFunc(provider func(context.Context) ([]cobra.Completion, error)) cobra.CompletionFunc {
	return func(cmd *cobra.Command, args []string, toComplete string) (completions []cobra.Completion, directive cobra.ShellCompDirective) {
		directive = cobra.ShellCompDirectiveNoFileComp
		if len(args) != 0 || provider == nil {
			return
		}
		// in case it gets stuck
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		// call provider
		completions, err := provider(ctx)
		if err != nil {
			return
		}
		return
	}
}
