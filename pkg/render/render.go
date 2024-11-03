package render

import (
	"bytes"
	"github.com/sangketkit01/go-web/pkg/config"
	"github.com/sangketkit01/go-web/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig
var functions = template.FuncMap{}

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get request template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatalln("template not found")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	if err := t.Execute(buf, td); err != nil {
		log.Fatalln(err)
	}

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Fatalln(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("../../templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("../../templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
