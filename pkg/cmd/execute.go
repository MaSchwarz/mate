package cmd

import (
	"log/slog"
	"mate/pkg/templates"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewExecute() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exec [template path]",
		Short: "Execute the given template",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			slog.Debug("start running exec command")

			openTemplate, err := templates.Load(args[0], cmd.InOrStdin())
			if err != nil {
				return err
			}

			sealedTemplate, err := openTemplate.Seal()
			if err != nil {
				return err
			}

			out := viper.GetString("out")
			if len(out) <= 0 {
				out = args[0] + ".output"
			}

			if err = sealedTemplate.Write(out); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringP("out", "o", "", "output file path")
	viper.BindPFlag("out", cmd.Flags().Lookup("out"))

	return cmd
}
