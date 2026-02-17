package generator

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Esabrina77/kago/internal/templates"
)

// AddFeature creates a new feature structure.
func AddFeature(projectDir, featureName string) error {
	// 0. Safety Check: Verify go.mod existence
	goModPath := filepath.Join(projectDir, "go.mod")
	if _, err := os.Stat(goModPath); os.IsNotExist(err) {
		return fmt.Errorf("go.mod not found in %s. \n❌ You must run this command from the root of a Go project", projectDir)
	}

	// 1. Feature Path: internal/features/<featureName>
	featurePath := filepath.Join(projectDir, "internal", "features", featureName)
	if err := os.MkdirAll(featurePath, 0755); err != nil {
		return err
	}

	// Basic capitalization: auth -> Auth
	capitalized := strings.ToUpper(featureName[:1]) + featureName[1:]

	// 4. Generate multiple files (Controller + Service + Repository)
	files := map[string]string{
		"module/controller.tmpl": fmt.Sprintf("%s_controller.go", featureName),
		"module/service.tmpl":    fmt.Sprintf("%s_service.go", featureName),
		"module/repository.tmpl": fmt.Sprintf("%s_repository.go", featureName),
	}

	for tmplName, fileName := range files {
		// Load Template
		tmplContent, err := fs.ReadFile(templates.GetTemplates(), tmplName)
		if err != nil {
			return fmt.Errorf("template %s not found: %v", tmplName, err)
		}

		// Parse
		tmpl, err := template.New(tmplName).Parse(string(tmplContent))
		if err != nil {
			return err
		}

		targetPath := filepath.Join(featurePath, fileName)
		if _, err := os.Stat(targetPath); !os.IsNotExist(err) {
			fmt.Printf("⚠️  Warning: %s already exists, skipping.\n", fileName)
			continue
		}

		f, err := os.Create(targetPath)
		if err != nil {
			return err
		}

		// Add FeatureLowercase for service struct naming
		type ExtendedData struct {
			FeatureName      string
			FeatureLowercase string
		}

		extData := ExtendedData{
			FeatureName:      capitalized,
			FeatureLowercase: strings.ToLower(capitalized),
		}

		if err := tmpl.Execute(f, extData); err != nil {
			f.Close()
			return err
		}
		f.Close()
		fmt.Printf("✅ Created: %s\n", targetPath)
	}

	return nil
}
