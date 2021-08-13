package handlers

import (
	"net/http"

	"github.com/tsawler/bookings/pkg/config"
	"github.com/tsawler/bookings/pkg/models"
	"github.com/tsawler/bookings/pkg/render"
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
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// println("entra about")
	stringMap := make(map[string]string)

	stringMap["test"] = "Hello, de nuevo"
	// println("entra about", stringMap["test"])

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
