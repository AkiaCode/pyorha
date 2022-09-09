package cli

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pyorha [options]",
	Short: "pyorha is serving tool",
}

func Run() {
	if err := rootCmd.Execute(); err != nil {
		errors.Wrap(err, "failed to execute root command")
	}
}
