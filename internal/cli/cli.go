package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/Esabrina77/kago/internal/generator"
)

// Execute runs the CLI logic.
func Execute() {
	// Flags
	projectType := flag.String("type", "simple", "Type of project (simple or web)")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		printUsage()
		os.Exit(1)
	}

	projectName := args[0]

	// Delegate to generator
	if err := generator.GenerateProject(projectName, *projectType); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: kago -type=<type> <project-name>")
	fmt.Println("\nAvailable types:")
	fmt.Println("  - simple (default)")
	fmt.Println("  - web")
}
