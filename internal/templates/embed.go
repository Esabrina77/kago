package templates

import (
	"embed"
	"io/fs"
)

// TemplatesEmbed is the embedded file system containing all templates.
//
//go:embed *
var templatesEmbed embed.FS

// GetTemplates returns the file system containing the templates.
func GetTemplates() fs.FS {
	return templatesEmbed
}
