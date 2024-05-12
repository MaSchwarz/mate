package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewVersion(version string, commit string, date string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "Version: %s, Commit Hash: %s, Build date: %s\n", version, commit, date)
		},
	}

	return cmd
}
