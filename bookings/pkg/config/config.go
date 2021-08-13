package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// Contiene la configuracion de la aplicacion
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
