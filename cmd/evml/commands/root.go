package commands

import (
	"github.com/Kdag-K/evm/cmd/evml/commands/run"
	"github.com/spf13/cobra"
)

// RootCmd is the root command for evml.
var RootCmd = &cobra.Command{ //nolint:exhaustivestruct
	Use:   "evml",
	Short: "EVM",
}

//nolint:gochecknoinits
func init() {
	RootCmd.AddCommand(
		run.RunCmd,
	)
	// do not print usage when error occurs.
	RootCmd.SilenceUsage = true
}
