package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed favicon.ico
var content embed.FS

//go:embed assets
var staticAssets embed.FS

// Static assets are mapped here.
type embeddedFileServer struct {
	root embed.FS
}

func MapStaticAssets() {
	if useEmbed {
		efs := &embeddedFileServer{staticAssets}
		
		// mux.Handle("/assets/", http.FileServer(http.FS(staticAssets)))
		mux.Handle("/assets/", efs.serveEmbedded(http.FileServer(http.FS(staticAssets))))
		mux.HandleFunc("/favicon.ico", faviconHandlerEmbeded)
	} else {
		mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
		mux.HandleFunc("/favicon.ico", faviconHandler)
	}

	log.Printf("Mapped static assets [embed: %t] \n", useEmbed)
}


func (efs *embeddedFileServer) serveEmbedded(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	    // Get file information
	    filePath := r.URL.Path
	    log.Print(filePath)
	    file, err := efs.root.Open(r.URL.Path)
	
	    if err != nil {
	        http.Error(w, err.Error(), http.StatusInternalServerError)
	        return
	    }
	
	    fileInfo, err := file.Stat()
	    if err != nil {
	        http.Error(w, err.Error(), http.StatusInternalServerError)
	        return
	    }
	
	    // Set custom Last-Modified header
	    lastModified := fileInfo.ModTime().UTC().Format(http.TimeFormat)
	    w.Header().Set("Last-Modified", lastModified)
	
	    // Use the default FileServer to serve the file
	    next.ServeHTTP(w, r)
	    file.Close()
    })
}

// Specific functions for working with root directory files
func faviconHandlerEmbeded(w http.ResponseWriter, r *http.Request) {
	favicon, err := content.ReadFile("favicon.ico")
	if err != nil {
		http.Error(w, "Favicon not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "image/x-icon")
	w.Write(favicon)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./favicon.ico")
}

// cache control for embedded files
// func cacheControl(h http.Handler) http.Handler {
// 	func(w http.ResponseWriter, r *http.Request) {
// 		h.ServeHTTP
// 	}
// }