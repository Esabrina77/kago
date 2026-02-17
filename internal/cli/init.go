package cli

import (
	"fmt"
	"os"

	"github.com/Esabrina77/kago/internal/generator"
	"github.com/spf13/cobra"
)

var projectType string

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

		fmt.Printf("Initialising %s project: %s\n", projectType, projectName)
		// Call our generator logic
		if err := generator.GenerateProject(projectName, projectType); err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("\n👉 Next steps:")
		fmt.Printf("  cd %s\n", projectName)
		fmt.Println("  kago add feature [name]")
	},
}

func init() {
	// Add init command to root
	rootCmd.AddCommand(initCmd)

	// Define flags for init command
	initCmd.Flags().StringVarP(&projectType, "type", "t", "simple", "Type of project (simple or web)")
}
