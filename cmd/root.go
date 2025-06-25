package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var flagJSON bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goignis",
	Short: "An optional CLI for ignis",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// GetRootCmd returns the root command instance.
func GetRootCmd() *cobra.Command {
	return rootCmd
}

// FlaggedJSON returns whether the flag "json" is set.
func FlaggedJSON() bool {
	return flagJSON
}

// RenderJSONIfFlagged prints the result in json if flag "json" is set, or do nothing and return false.
func RenderJSONIfFlagged(result any) bool {
	if flagJSON {
		data, _ := json.Marshal(result)
		fmt.Printf("%s\n", data)
	}
	return flagJSON
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&flagJSON, "json", "j", false, "Print results in json")
}
