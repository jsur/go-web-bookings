package handlers

import (
	"net/http"

	"github.com/jsur/go-web-bookings/pkg/config"
	"github.com/jsur/go-web-bookings/pkg/models"
	"github.com/jsur/go-web-bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.GetTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.GetTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.GetTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Generals renders generals page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.GetTemplate(w, "generals.page.tmpl", &models.TemplateData{})
}

// Majors renders majors page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.GetTemplate(w, "majors.page.tmpl", &models.TemplateData{})
}

// Availability renders availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.GetTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}

// Contact renders contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.GetTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}
