package pages

import (
	"net/http"
)

var fs http.Handler

func Init() {
	fs = http.FileServer(http.Dir("wwwroot"))
}
