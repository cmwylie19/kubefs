package main

import (
	"github.com/spf13/cobra"
)

var Verbose bool
var baseCmd = &cobra.Command{
	Use:   "kubefs",
	Short: "Standalone or Kubernetes filesystem controller",
}

// GetRootCommand returns the root cobra command to be executed
// by main.
func GetRootCommand() *cobra.Command {
	baseCmd.AddCommand(getServerCommand())
	// baseCmd.AddCommand(getServerCommand())
	return baseCmd
}
