package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// ExampleGetRootCmd shows adding custom subcommands to root command.
func ExampleGetRootCmd() {
	var flagVerbose bool
	// create a subcommand
	fooCmd := cobra.Command{
		Use: "ping",
		Run: func(cmd *cobra.Command, _ []string) {
			if flagVerbose {
				cmd.Println("Verbose!")
			}
			fmt.Println("Pong!")
		},
	}
	fooCmd.Flags().BoolVar(&flagVerbose, "verbose", false, "Verbose output")

	// add subcommand to root
	rootCmd := GetRootCmd()
	rootCmd.AddCommand(&fooCmd)

	// override command args for testing
	rootCmd.SetArgs([]string{"ping", "--verbose"})
	err := rootCmd.Execute()
	if err != nil {
		log.Println("Failed:", err)
		return
	}

	// Output: Pong!
}
