package handlers

import (
	"fmt"
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
	render.GetTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.GetTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.GetTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Generals renders generals page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.GetTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// Majors renders majors page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.GetTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Availability renders availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.GetTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability handles new availability POST calls
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

// Contact renders contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.GetTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}
