package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var build = "-1"
var commit = "000000000"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cdr",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`{"build":"%s","commit":"%s"}`, build, commit)
	},
}
