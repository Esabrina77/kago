package cli

import (
	"fmt"
	"os"

	"github.com/Esabrina77/kago/internal/generator"
	"github.com/spf13/cobra"
)

var (
	// Flags
	projectType string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kago",
	Short: "kaGO is a professional Go project bootstrapper",
	Long: `kaGO helps you bootstrap Go applications with a clean architecture.
It allows you to initialize projects and generate code components.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new Go project",
	Long: `Initialize a new Go project with the specified structure.
Example:
  kago init my-awesome-project --type=web`,
	Args: cobra.ExactArgs(1), // We explicitly require 1 argument: project-name
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		// Call our generator logic
		if err := generator.GenerateProject(projectName, projectType); err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Add init command to root
	rootCmd.AddCommand(initCmd)

	// Define flags for init command
	initCmd.Flags().StringVarP(&projectType, "type", "t", "simple", "Type of project (simple or web)")
}
