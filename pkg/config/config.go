package config

import "html/template"

// AppConfig holds application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
