package handlers

import (
	"changeme/assets"
	"html/template"
)

// Full page templates
var indexTemplate *template.Template

type templateData struct {
	SearchTerm string
}

func (app *Application) ParseTemplates() error {
	t, err := template.ParseFS(assets.Assets, "public/html/index.html", "public/html/navbar.tmpl", "public/html/footer.tmpl")
	if err != nil {
		return err
	}
	indexTemplate = t

	return nil
}
