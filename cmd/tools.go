package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(toolsCmd)
}

var toolsCmd = &cobra.Command{
	Use:   "tools",
	Short: "Show the Opsta version and information",
	Long:  "Show the Opsta version and information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("opsta-cli version 0.0.1")
	},
}
