package pages

import "net/http"

var HttpFS  http.Handler

func Init() {
	HttpFS = http.FileServer(http.Dir("wwwroot"))
}