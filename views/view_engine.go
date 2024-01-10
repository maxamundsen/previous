package views

import (
	"gohttp/constants/build"
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
	
	if build.DEVEL {
		templates, err = parseTemplates()
	} else {
		once.Do(func() {
			templates, err = parseTemplates()
		})
	}
	
	return templates, err
}

func parseTemplates() (*template.Template, error) {

	var err error

	if build.DEVEL {
		templates, err = template.ParseGlob("views/*.html")
	} else {
		templates, err = template.ParseFS(embeddedTemplates, "*.html")
	}

	
	if err != nil {
		return nil, err
	}
	
	return templates, nil
}

