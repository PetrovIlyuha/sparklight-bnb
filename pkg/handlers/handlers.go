package handlers

import (
	"net/http"

	"github.com/PetrovIlyuha/sparklight-bnb/pkg/config"
	"github.com/PetrovIlyuha/sparklight-bnb/pkg/models"
	"github.com/PetrovIlyuha/sparklight-bnb/pkg/render"
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

func (repo *Repository) HomePage(rw http.ResponseWriter, req *http.Request) {
	remoteIP := req.RemoteAddr
	repo.App.Session.Put(req.Context(), "remote_ip", remoteIP)
	render.HydrateTemplate(rw, "home.page.tmpl", &models.TemplateData{})
}

func (repo *Repository) AboutPage(rw http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	stringMap["greeting"] = "Hello, Dear User! 😷"
	remoteIp := repo.App.Session.GetString(req.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp
	render.HydrateTemplate(rw, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
