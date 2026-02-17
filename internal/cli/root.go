package cli

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kago",
	Short: "kaGO is a professional Go project bootstrapper",
	Long: `kaGO helps you bootstrap Go applications with a clean architecture.
It allows you to initialize projects and generate code components.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
