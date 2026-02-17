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
	// 1. Feature Path: internal/features/<featureName>
	featurePath := filepath.Join(projectDir, "internal", "features", featureName)
	if err := os.MkdirAll(featurePath, 0755); err != nil {
		return err
	}

	// 2. Load Template
	tmplContent, err := fs.ReadFile(templates.GetTemplates(), "feature.tmpl")
	if err != nil {
		return fmt.Errorf("feature template not found: %v", err)
	}

	// 3. Prepare Template Config
	tmpl, err := template.New("feature").Parse(string(tmplContent))
	if err != nil {
		return err
	}

	type FeatureData struct {
		PackageName string
		FeatureName string
	}

	data := FeatureData{
		PackageName: featureName,
		FeatureName: strings.Title(featureName), // e.g. "auth" -> "Auth"
	}

	// 4. Create File (e.g. auth_handler.go)
	filePath := filepath.Join(featurePath, fmt.Sprintf("%s.go", featureName))
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := tmpl.Execute(file, data); err != nil {
		return err
	}

	fmt.Printf("Created file: %s\n", filePath)
	return nil
}
