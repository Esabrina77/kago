package generator

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Esabrina77/kago/internal/templates"
)

type LayerDefinition struct {
	Directory string // e.g. "controllers"
	Template  string // e.g. "module/controller.tmpl"
	Suffix    string // e.g. "_controller.go"
}

// AddFeature creates a new feature using Layered Architecture.
func AddFeature(projectDir, featureName string) error {
	// 0. Safety Check & Get Module Name
	moduleName, err := getModuleName(projectDir)
	if err != nil {
		return err
	}

	if len(featureName) == 0 {
		return fmt.Errorf("feature name cannot be empty")
	}
	// Capitalize: users -> Users
	capitalized := strings.ToUpper(featureName[:1]) + featureName[1:]

	// 1. Define Layers
	layers := []LayerDefinition{
		{Directory: "controllers", Template: "module/controller.tmpl", Suffix: "_controller.go"},
		{Directory: "services", Template: "module/service.tmpl", Suffix: "_service.go"},
		{Directory: "repositories", Template: "module/repository.tmpl", Suffix: "_repository.go"},
	}

	for _, layer := range layers {
		// Target Directory: internal/controllers
		layerDir := filepath.Join(projectDir, "internal", layer.Directory)
		if err := os.MkdirAll(layerDir, 0755); err != nil {
			return err
		}

		// Target File: users_controller.go
		fileName := fmt.Sprintf("%s%s", strings.ToLower(featureName), layer.Suffix)
		targetPath := filepath.Join(layerDir, fileName)

		// Check existence
		if _, err := os.Stat(targetPath); !os.IsNotExist(err) {
			fmt.Printf("⚠️  Warning: %s already exists, skipping.\n", fileName)
			continue
		}

		// Load Template
		tmplContent, err := fs.ReadFile(templates.GetTemplates(), layer.Template)
		if err != nil {
			return fmt.Errorf("template %s not found: %v", layer.Template, err)
		}

		// Parse
		tmpl, err := template.New(layer.Template).Parse(string(tmplContent))
		if err != nil {
			return err
		}

		// Data for template
		type TemplateData struct {
			FeatureName      string // Users
			FeatureLowercase string // users
			ModulePath       string // github.com/user/project
		}

		data := TemplateData{
			FeatureName:      capitalized,
			FeatureLowercase: strings.ToLower(capitalized),
			ModulePath:       moduleName,
		}

		// Create File
		f, err := os.Create(targetPath)
		if err != nil {
			return err
		}

		if err := tmpl.Execute(f, data); err != nil {
			f.Close()
			return err
		}
		f.Close()
		fmt.Printf("✅ Created: %s\n", targetPath)
	}

	return nil
}

// getModuleName extracts the module directive from go.mod
func getModuleName(projectDir string) (string, error) {
	goModPath := filepath.Join(projectDir, "go.mod")
	file, err := os.Open(goModPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("go.mod not found in %s.\n❌ You must run this command from the root of a Go project", projectDir)
		}
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Read line by line until "module" keyword is found
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				return parts[1], nil
			}
		}
	}

	return "", fmt.Errorf("could not find 'module' definition in go.mod")
}
