package handlers

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// set timestamp at initial runtime for embedded caching
var currentTime time.Time

// helper function to map endpoints for static assets.
func MapStaticAssets(mux *http.ServeMux) {
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("wwwroot/assets"))))
	mux.Handle("/favicon.ico", http.FileServer(http.Dir("wwwroot")))
	log.Println("Mapped static assets [embed: false]")
}

// helper function to map endpoints for embedded static assets.
func MapStaticAssetsEmbed(mux *http.ServeMux, fs *embed.FS) {
	currentTime = time.Now()
	mux.Handle("/assets/", serveEmbedded(fs))
	mux.Handle("/favicon.ico", serveEmbedded(fs))
	log.Println("Mapped static assets [embed: true]")
}

// When static assets are served out of the embedded file system, you need to append
// the proper http headers to allow caching. by default, the embedded fs has no 'last modified'
// property, so you must add one, and specify that you want cache control on.
func serveEmbedded(fs *embed.FS) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		location := "wwwroot" + r.URL.Path

		file, err := fs.ReadFile(location)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		var contentType string

		if strings.HasSuffix(location, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(location, ".js") {
			contentType = "text/javascript"
		} else {
			contentType = http.DetectContentType(file)
		}

		lastModified := currentTime.UTC().Format(http.TimeFormat)

		w.Header().Set("Cache-Control", "max-age=604800")
		w.Header().Set("Last-Modified", lastModified)
		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(file)))

		w.Write(file)
	})
}
