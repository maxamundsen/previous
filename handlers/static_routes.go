package handlers

import (
	"gohttp/constants"
	"embed"
	"log"
	"net/http"
)

// Static assets are mapped here.
type embeddedFileServer struct {
	root embed.FS
}

func MapStaticAssetsEmbed(mux *http.ServeMux, fs *embed.FS) {
	mux.Handle("/assets/", serveEmbedded(http.FileServer(http.FS(fs)), fs))
	log.Printf("Mapped static assets [embed: %t] \n", constants.UseEmbed)
}

func MapStaticAssets(mux *http.ServeMux) {
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("wwwroot/assets"))))
	log.Printf("Mapped static assets [embed: %t] \n", constants.UseEmbed)
}

func MapFavicon(mux *http.ServeMux, fs *embed.FS) {
	if constants.UseEmbed {
		mux.Handle("/favicon.ico", faviconHandlerEmbedded(fs))	
	} else {
		mux.HandleFunc("/favicon.ico", faviconHandler)		
	}
}

func serveEmbedded(next http.Handler, fs *embed.FS) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := fs.Open(r.URL.Path)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fileInfo, err := file.Stat()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		lastModified := fileInfo.ModTime().UTC().Format(http.TimeFormat)
		w.Header().Set("Last-Modified", lastModified)

		next.ServeHTTP(w, r)
		file.Close()
	})
}

// Specific functions for working with root directory files
func faviconHandlerEmbedded(fs *embed.FS) http.Handler {
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

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../wwwroot/favicon.ico")
}

// cache control for embedded files
// func cacheControl(h http.Handler) http.Handler {
// 	func(w http.ResponseWriter, r *http.Request) {
// 		h.ServeHTTP
// 	}
// }
