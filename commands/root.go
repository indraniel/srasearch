package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var version string

var rootCmd = &cobra.Command{
	Use:   "srasearch",
	Short: "A SRA Search Utility",
	Long:  `A utility to manage and view local SRA submissions data`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s -- see '%s help'\n",
			"Please specify a subcommand", os.Args[0])
	},
}

func Execute(semanticVersion string) {
	version = semanticVersion
	AddCommands()
	rootCmd.Execute()
}

func AddCommands() {
	rootCmd.AddCommand(cmdInitDump)
	rootCmd.AddCommand(cmdVersion)
	rootCmd.AddCommand(cmdMakeIndex)
	rootCmd.AddCommand(cmdIncrementDump)
}
