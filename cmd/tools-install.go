package cmd

import (
	"github.com/opsta/opsta-cli/internal/text"
	"github.com/spf13/cobra"
)

func init() {
	toolsCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "The easy way to install the most DevOps in a single command",
	Long:  "The easy way to install the most DevOps in a single command" + text.Logo(),
	Run: func(cmd *cobra.Command, args []string) {

	},
}
