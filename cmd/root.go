// Package cmd ...
package cmd

import (
	"github.com/opsta/opsta-cli/internal/text"
	"github.com/spf13/cobra"
)

// RootCmd is the root cmd in the cli use for travel to subcommand
var RootCmd = &cobra.Command{
	Use:   "opsta",
	Short: "opsta is a automate tool to automate devops in the simple way 😀",
	Long:  "A DevOps Tools Manager to automate thing in The simple way 😀" + text.Logo(),
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
