package handlers

import (
	"github.com/sangketkit01/go-web/pkg/config"
	"github.com/sangketkit01/go-web/pkg/models"
	"github.com/sangketkit01/go-web/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	intMap := make(map[string]int)
	intMap["int1"] = 10

	model := &models.TemplateData{
		StringMap: stringMap,
		IntMap:    intMap,
	}

	render.RenderTemplate(w, "about.page.tmpl", model)
}
