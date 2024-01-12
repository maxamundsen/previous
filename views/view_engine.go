package views

import (
	"gohttp/auth"
	"gohttp/constants/build"
	"html/template"
	"net/http"
	"sync"
)

var (
	templates *template.Template
	once      sync.Once
)

type ViewBase struct {
	Identity *auth.Identity
	ViewData map[string]interface{}
}

func NewViewBase(user *auth.Identity, viewData map[string]interface{}) ViewBase {
	viewBase := ViewBase{
		user,
		viewData,
	}

	return viewBase
}

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
