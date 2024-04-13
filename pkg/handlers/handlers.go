package handlers

import (
	"myapp/pkg/config"
	"myapp/pkg/models"
	"myapp/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (repository *Repository) Home(w http.ResponseWriter, r *http.Request) {

	var StringMap = make(map[string]string)
	StringMap["test"] = "Hello, again. data from the map"

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: StringMap,
	})
}

func (repository *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}
