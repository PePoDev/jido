package cmd

import (
	"github.com/opsta/opsta-cli/internal/text"
	"github.com/spf13/cobra"
)

func init() {
	toolsCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove tool that install by Opsta-Cli",
	Long:  "Remove tool that install by Opsta-Cli" + text.Logo(),
	Run: func(cmd *cobra.Command, args []string) {

	},
}
