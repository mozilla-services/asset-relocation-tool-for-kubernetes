package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const AppName = "relok8s"

var rootCmd = &cobra.Command{
	Use: AppName,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// rootCmd may return an error, but Cobra is already displaying it
		// so here we just swallow but still exit with an error code
		os.Exit(1)
	}
}
