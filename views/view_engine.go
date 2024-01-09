package views

import (
	"gohttp/constants"
	"html/template"
	"net/http"
	"sync"
)

var (
	templates         *template.Template
	once              sync.Once
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmpls, err := loadTemplates()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpls.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loadTemplates() (*template.Template, error) {
	var err error
	
	once.Do(func() {
		templates, err = parseTemplates()
	})
	
	return templates, err
}

func parseTemplates() (*template.Template, error) {

	var err error

	if constants.EMBED {
		templates, err = template.ParseFS(embeddedTemplates, "*.html")
	} else {
		templates, err = template.ParseGlob("views/*.html")
	}

	
	if err != nil {
		return nil, err
	}
	
	return templates, nil
}

