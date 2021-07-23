package config

import (
	"html/template"
	"log"
)

// Contiene la configuracion de la aplicacion
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
