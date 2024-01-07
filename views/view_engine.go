package views

import (
	"gohttp/constants"
	"embed"
	"html/template"
	"net/http"
	"sync"
	// "errors"
)

var (
	//go:embed *.html
	embeddedTemplates embed.FS
	templates         *template.Template
	once              sync.Once
)

func parseTemplates() (*template.Template, error) {
	var err error

	if constants.UseEmbed {
		templates, err = template.ParseFS(embeddedTemplates, "*.html")
	} else {
		templates, err = template.ParseGlob("views/*.html")
	}

	if err != nil {
		return nil, err
	}
	return templates, nil
}

func loadTemplates() (*template.Template, error) {
	var err error
	
	// once.Do(func() {
		templates, err = parseTemplates()
	// })
	
	// if templates == nil {
	// 	err = errors.New("Templates is null")
	// }
	
	return templates, err
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
