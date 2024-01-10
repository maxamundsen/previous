package handlers

import (
	"time"
	"strings"
	"fmt"
	"embed"
	"log"
	"net/http"
)

// set timestamp at initial runtime
var currentTime time.Time

func MapStaticAssets(mux *http.ServeMux) {
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("wwwroot/assets"))))
	mux.Handle("/favicon.ico", http.FileServer(http.Dir("wwwroot")))
	log.Println("Mapped static assets [embed: false]")
}

func MapStaticAssetsEmbed(mux *http.ServeMux, fs *embed.FS) {
	currentTime = time.Now()
	mux.Handle("/assets/", serveEmbedded(fs))
	mux.Handle("/favicon.ico", serveEmbedded(fs))
	log.Println("Mapped static assets [embed: true]")
}

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
