package views

import (
	"gohttp/auth"
	"gohttp/build"
	"html/template"
	"net/http"
	"sync"
)

// The view engine is a wrapper around the standard html/template functions.
// A ViewModel struct is provided, which is passed to a view template

// The ViewBase contains a map with string keys, interface{} values.
// This means that you can add any data of any type you want to the ViewData map.
// Because of this structure, you do not need to create page specific models, but instead
// load the map with whatever data you need, and retrieve it. You can still perform extremely simple
// logic on map values such as 'does this exist' or 'is this a value'

var (
	templates *template.Template
	once      sync.Once
)

type ViewModel struct {
	Identity *auth.Identity
	ViewData map[string]interface{}
}

func NewViewModel(user *auth.Identity, viewData map[string]interface{}) ViewModel {
	model := ViewModel{
		user,
		viewData,
	}

	return model
}

// RenderTemplate executes the template with the provided data.
// Unlike the default, this wrapper forces the input data
// to be a ViewModel struct. This is because a ViewModel struct contains
// a map[string]interface{} where any data type can be passed.
func RenderTemplate(w http.ResponseWriter, tmpl string, model ViewModel) {
	tmpls, err := loadTemplates()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpls.ExecuteTemplate(w, tmpl+".html", model)
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
