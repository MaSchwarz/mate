package cmd

import (
	"github.com/spf13/cobra"
)

func NewRoot(version string, commit string, date string) *cobra.Command {
	cmd := &cobra.Command{
		Use: "mate",
	}

	cmd.AddCommand(NewVersion(version, commit, date))
	cmd.AddCommand(NewExecute())

	return cmd
}
