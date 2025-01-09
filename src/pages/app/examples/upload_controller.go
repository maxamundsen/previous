package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "previous/pages/app"

	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"previous/middleware"
	"previous/auth"
)

// @Identity
// @Protected
// @CookieSession
func UploadController(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	if r.Method == http.MethodPost {
		r.ParseMultipartForm(10 << 20)

		file, fileHeader, err := r.FormFile("file")

		if err != nil {
			UploadView(err.Error(), "", *identity).Render(w)
			return
		}

		defer file.Close()

		err = os.MkdirAll("./uploads", os.ModePerm)
		if err != nil {
			UploadView(err.Error(), "", *identity).Render(w)
			return
		}

		dst, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
		if err != nil {
			UploadView(err.Error(), "", *identity).Render(w)
			return
		}

		defer dst.Close()

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			UploadView(err.Error(), "", *identity).Render(w)
			return
		}

		dst.Write(fileBytes)

		UploadView("", "Successfully uploaded file", *identity).Render(w)
	} else {
		UploadView("", "", *identity).Render(w)
	}
}

func UploadView(errorMsg string, successMsg string, identity auth.Identity) Node {
	return AppLayout("Upload Example", identity,
		If(errorMsg != "", Div(Class("alert alert-danger"), Text(errorMsg))),
		If(successMsg != "", Div(Class("alert alert-success"), Text(successMsg))),
		Form(Action("/app/examples/upload"), Method("post"), EncType("multipart/form-data"),
			Input(Type("file"), Name("file"), Class("form-control")),
			Br(),
			Button(Type("submit"), Class("btn btn-primary"), Text("Upload")),
		),
	)
}
