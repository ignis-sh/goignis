package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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

// renderJSON returns whether the flag 'json' is set
func renderJSON() bool {
	useJSON, _ := rootCmd.PersistentFlags().GetBool("json")
	return useJSON
}

// renderResultIfJSON prints the result in json if flag 'json' is set, or do nothing and return false
func renderResultIfJSON(result any) bool {
	useJSON := renderJSON()
	if useJSON {
		data, _ := json.Marshal(result)
		fmt.Printf("%s\n", data)
	}
	return useJSON
}

func init() {
	rootCmd.PersistentFlags().BoolP("json", "j", false, "Print results in json")
}
