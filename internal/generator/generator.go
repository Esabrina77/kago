package generator

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Esabrina77/kago/internal/templates"
)

// GenerateProject creates a new Go project structure.
func GenerateProject(projectName, projectType string) error {
	// 1. Validation
	if err := checkDir(projectName); err != nil {
		return err
	}

	fmt.Printf("🚀 Initializing %s project: %s...\n", projectType, projectName)

	// 2. Scaffolding
	switch projectType {
	case "web":
		if err := createWebStructure(projectName); err != nil {
			return err
		}
	case "simple":
		if err := createSimpleStructure(projectName); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown project type: %s", projectType)
	}

	// 3. Module Init
	if err := initModule(projectName); err != nil {
		fmt.Printf("⚠️  Warning: could not initialize go module: %v\n", err)
	}

	fmt.Println("✅ Done! Your project is ready.")
	return nil
}

func checkDir(name string) error {
	if _, err := os.Stat(name); !os.IsNotExist(err) {
		return fmt.Errorf("directory '%s' already exists", name)
	}
	return nil
}

func createSimpleStructure(projectName string) error {
	if err := os.MkdirAll(projectName, 0755); err != nil {
		return err
	}
	return generateFile(projectName, "main.go", "project/simple/main.go.tmpl")
}

func createWebStructure(projectName string) error {
	subfolders := []string{
		"cmd",
		"internal",
		"pkg",
		"api",
	}

	for _, sub := range subfolders {
		targetPath := filepath.Join(projectName, sub)
		if err := os.MkdirAll(targetPath, 0755); err != nil {
			return err
		}
	}

	// Create cmd/main.go
	if err := os.MkdirAll(filepath.Join(projectName, "cmd"), 0755); err != nil {
		return err
	}

	// We map the template "web_main.tmpl" to "cmd/main.go"
	// But generateFile takes relative path from project root? Yes.
	return generateFile(projectName, filepath.Join("cmd", "main.go"), "project/web/main.go.tmpl")
}

func generateFile(projectDir, deployPath, templateName string) error {
	// Read from embedded FS
	tmplContent, err := fs.ReadFile(templates.GetTemplates(), templateName)
	if err != nil {
		return fmt.Errorf("template not found: %s", templateName)
	}

	// Replace placeholders if any (currently manual replacement for simplicity)
	content := string(tmplContent)
	// Example: content = strings.ReplaceAll(content, "{{ProjectName}}", projectDir)
	// (We can add this later, for now just copy)

	fullPath := filepath.Join(projectDir, deployPath)
	return os.WriteFile(fullPath, []byte(content), 0644)
}

func initModule(projectName string) error {
	// Extract module name from project name usually, assuming github.com/user/project
	// For now, keep it simple "go mod init <projectName>"

	// If project name is a path (e.g. github.com/foo/bar), use it as module name, but folder is last part?
	// Existing logic was: go mod init <projectName> inside <projectName> folder.
	// We assume projectName is just the folder name for now.

	moduleName := projectName
	if strings.Contains(projectName, "/") || strings.Contains(projectName, "\\") {
		// Basic sanitization
		parts := strings.Split(filepath.ToSlash(projectName), "/")
		moduleName = parts[len(parts)-1]
	}

	fmt.Printf("📦 Running 'go mod init %s'...\n", moduleName)
	cmd := exec.Command("go", "mod", "init", moduleName)
	cmd.Dir = projectName
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
