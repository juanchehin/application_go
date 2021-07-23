package handlers

import (
	"net/http"

	"github.com/tsawler/go-course/pkg/config"
	"github.com/tsawler/go-course/pkg/models"
	"github.com/tsawler/go-course/pkg/render"
)

// Repositorio usado para los controladores (handlers)
var Repo *Repository

// Tipo de repositorio
type Repository struct {
	App *config.AppConfig
}

// Crea un nuevo repositorio
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Establece el repositorio para el controlador
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// println("entra about")
	stringMap := make(map[string]string)

	stringMap["test"] = "Hello, de nuevo"
	println("entra about", stringMap["test"])

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}
