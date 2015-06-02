package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Display version information",
	Long:  `Display the version details`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
