package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/lost-melody/goignis/pkg"

	"github.com/spf13/cobra"
)

func init() {
	// subcommands without args or return values
	addPlainCmd("quit", "Quit Ignis", pkg.QuitIgnis)
	addPlainCmd("reload", "Reload Ignis", pkg.ReloadIgnis)
	addPlainCmd("inspector", "Open GTK Inspector", pkg.OpenInspector)
	// subcommands with no args but one single return value
	addPlainCmdWithResult("list-windows", "List names of all windows", pkg.ListWindows, func(windows []string) {
		fmt.Println(strings.Join(windows, "\n"))
	})
	// subcommands with one single argument and one single return value
	addPlainCmdWithArgResult("toggle-window", "Toggle a window", pkg.ToggleWindow, pkg.ListWindows, func(found bool) {
		fmt.Println(found)
	})
	addPlainCmdWithArgResult("open-window", "Open a window", pkg.OpenWindow, pkg.ListWindows, func(found bool) {
		fmt.Println(found)
	})
	addPlainCmdWithArgResult("close-window", "Close a window", pkg.CloseWindow, pkg.ListWindows, func(found bool) {
		fmt.Println(found)
	})
	// For more complicated subcommands, use `cobra-cli add xxx` instead
}

func addPlainCmd(use, short string, callback func(context.Context) error) {
	rootCmd.AddCommand(&cobra.Command{
		Use:               use,
		Short:             short,
		Args:              cobra.NoArgs,
		ValidArgsFunction: argsCompletionFunc(nil),
		Run: func(cmd *cobra.Command, args []string) {
			err := callback(context.Background())
			if err != nil {
				cmd.PrintErrln("Failed:", err)
				return
			}
		},
	})
}

func addPlainCmdWithArg(
	use, short string,
	callback func(context.Context, string) error,
	completion func(context.Context) ([]cobra.Completion, error),
) {
	rootCmd.AddCommand(&cobra.Command{
		Use:               use,
		Short:             short,
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: argsCompletionFunc(completion),
		Run: func(cmd *cobra.Command, args []string) {
			err := callback(context.Background(), args[0])
			if err != nil {
				cmd.PrintErrln("Failed:", err)
				return
			}
		},
	})
}

func addPlainCmdWithResult[T any](
	use, short string,
	callback func(context.Context) (T, error),
	render func(T),
) {
	rootCmd.AddCommand(&cobra.Command{
		Use:               use,
		Short:             short,
		Args:              cobra.NoArgs,
		ValidArgsFunction: argsCompletionFunc(nil),
		Run: func(cmd *cobra.Command, args []string) {
			result, err := callback(context.Background())
			if err != nil {
				cmd.PrintErrln("Failed:", err)
				return
			}
			if !renderResultIfJSON(result) {
				render(result)
			}
		},
	})
}

func addPlainCmdWithArgResult[T any](
	use, short string,
	callback func(context.Context, string) (T, error),
	completion func(context.Context) ([]cobra.Completion, error),
	render func(T),
) {
	rootCmd.AddCommand(&cobra.Command{
		Use:               use,
		Short:             short,
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: argsCompletionFunc(completion),
		Run: func(cmd *cobra.Command, args []string) {
			result, err := callback(context.Background(), args[0])
			if err != nil {
				cmd.PrintErrln("Failed:", err)
				return
			}
			if !renderResultIfJSON(result) {
				render(result)
			}
		},
	})
}

func argsCompletionFunc(provider func(context.Context) ([]cobra.Completion, error)) cobra.CompletionFunc {
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
