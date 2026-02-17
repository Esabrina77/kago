package cli

import (
	"fmt"
	"os"

	"github.com/Esabrina77/kago/internal/generator"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [resource]",
	Short: "Add a resource to your project",
	Long:  `Add a resource like a feature, controller, or service to your existing project.`,
}

var addFeatureCmd = &cobra.Command{
	Use:   "feature [name]",
	Short: "Add a new feature module",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		featureName := args[0]
		projectDir, _ := os.Getwd() // Assume current directory is project root

		fmt.Printf("🏗️  Adding feature '%s' to project...\n", featureName)
		if err := generator.AddFeature(projectDir, featureName); err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("✅ Feature '%s' added successfully!\n", featureName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.AddCommand(addFeatureCmd)
}
