package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// 1. TEMPLATES EMBEDDING

//go:embed templates/simple.tmpl
var simpleTemplate string

//go:embed templates/web_main.tmpl
var webMainTemplate string

func main() {
	// 2. FLAG SETUP
	projectType := flag.String("type", "simple", "Type of project (simple or web)")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: kago -type=<type> <project-name>")
		os.Exit(1)
	}
	projectName := args[0]

	// 3. VALIDATIONS
	validateType(*projectType)
	checkDir(projectName)

	// 4. EXECUTION
	fmt.Printf("🚀 Initializing %s project: %s...\n", *projectType, projectName)

	switch *projectType {
	case "web":
		createWebStructure(projectName)
	case "simple":
		createSimpleStructure(projectName)
	}

	// 5. AUTO-MODULE INIT (Phase 3)
	initModule(projectName)

	fmt.Println("✅ Done! Your project is ready.")
}

// --- HELPER FUNCTIONS ---

func validateType(t string) {
	if t != "simple" && t != "web" {
		fmt.Printf("Error: unknown type '%s'. Use 'simple' or 'web'.\n", t)
		os.Exit(1)
	}
}

func checkDir(name string) {
	if _, err := os.Stat(name); !os.IsNotExist(err) {
		fmt.Printf("Error: directory '%s' already exists\n", name)
		os.Exit(1)
	}
}

func createSimpleStructure(projectName string) {
	os.MkdirAll(projectName, 0755)
	generateFile(filepath.Join(projectName, "main.go"), simpleTemplate)
}

func createWebStructure(projectName string) {
	// Create folders
	subfolders := []string{
		"cmd",
		"internal",
		"pkg",
		"api",
	}

	for _, sub := range subfolders {
		targetPath := filepath.Join(projectName, sub)
		os.MkdirAll(targetPath, 0755)
	}

	// Generate main.go in cmd/
	mainPath := filepath.Join(projectName, "cmd", "main.go")
	generateFile(mainPath, webMainTemplate)
}

func generateFile(path string, content string) {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error creating file at %s: %v\n", path, err)
		os.Exit(1)
	}
}

func initModule(projectName string) {
	fmt.Printf("📦 Running 'go mod init %s'...\n", projectName)
	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = projectName
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("⚠️  Warning: could not initialize go module: %v\n", err)
	}
}
