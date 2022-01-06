package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/PetrovIlyuha/sparklight-bnb/pkg/config"
	"github.com/PetrovIlyuha/sparklight-bnb/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AppDefaultData(templateData *models.TemplateData) *models.TemplateData {
	return templateData
}

func HydrateTemplate(rw http.ResponseWriter, tmpl string, templateData *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	pageTemplate, ok := tc[tmpl]
	if !ok {
		log.Fatal("Oops. Could not get templates from templates cache")
	}
	buf := new(bytes.Buffer)
	templateData = AppDefaultData(templateData)
	_ = pageTemplate.Execute(buf, templateData)
	_, err := buf.WriteTo(rw)
	if err != nil {
		fmt.Println("Server crashed while parsing templates!")
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {

		pagename := filepath.Base(page)

		ts, err := template.New(pagename).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[pagename] = ts
	}
	return myCache, nil
}
