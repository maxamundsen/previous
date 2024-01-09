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

// Static assets are mapped here.
type embeddedFileServer struct {
	root embed.FS
}

func MapStaticAssets(mux *http.ServeMux) {
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("wwwroot/assets"))))
	log.Println("Mapped static assets [embed: false]")
}

func MapStaticAssetsEmbed(mux *http.ServeMux, fs *embed.FS) {
	currentTime = time.Now()
	mux.Handle("/assets/", serveEmbedded(fs))
	log.Println("Mapped static assets [embed: true]")
}

func MapFavicon(mux *http.ServeMux, fs *embed.FS) {
	mux.HandleFunc("/favicon.ico", faviconHandler)
}

// map favicon
func MapFaviconEmbed(mux *http.ServeMux, fs *embed.FS) {
	mux.Handle("/favicon.ico", faviconHandlerEmbed(fs))
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../wwwroot/favicon.ico")
}

// Specific functions for working with root directory files
func faviconHandlerEmbed(fs *embed.FS) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		favicon, err := fs.ReadFile("favicon.ico")
		if err != nil {
			http.Error(w, "Favicon not found", http.StatusNotFound)
			return
		}
	
		w.Header().Set("Content-Type", "image/x-icon")
		w.Write(favicon)
	
	})
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
		
		if (strings.HasSuffix(location, ".css")) {
			contentType = "text/css"
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
