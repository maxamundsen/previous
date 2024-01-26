package handlers

import (
	"gohttp/auth"
	"gohttp/data"
	"gohttp/views"
	"net/http"
	"net/mail"
	"strconv"
	"log"
	"io/ioutil"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	identity := sessionStore.GetIdentityFromCtx(r)

	viewData := make(map[string]interface{})
	viewData["Title"] = "Example Page"

	model := views.NewViewModel(identity, viewData)
	views.RenderTemplate(w, "example", model)
}

func exampleCounterHandler(w http.ResponseWriter, r *http.Request) {
	viewData := make(map[string]interface{})

	formVal := r.FormValue("number")

	num, _ := strconv.Atoi(formVal)

	viewData["Number"] = num + 1

	model := views.NewViewModel(nil, viewData)
	views.RenderTemplate(w, "example_counter", model)
}

func examplePassgenHandler(w http.ResponseWriter, r *http.Request) {
	identity := sessionStore.GetIdentityFromCtx(r)

	if identity.Claims["CanGeneratePasswords"] != "true" {
		return
	}

	val1 := r.FormValue("val1")

	var password string

	if val1 == "" {
		password = "empty"
	} else {
		password, _ = auth.HashPassword(val1)
	}

	viewData := make(map[string]interface{})
	viewData["Password"] = password

	model := views.NewViewModel(identity, viewData)

	views.RenderTemplate(w, "example_passgen", model)
}

func exampleDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	identity := sessionStore.GetIdentityFromCtx(r)

	viewData := make(map[string]interface{})
	viewData["Title"] = "Database Example"

	model := views.NewViewModel(identity, viewData)

	views.RenderTemplate(w, "example_database", model)
}

func exampleAdduserHandler(w http.ResponseWriter, r *http.Request) {
	viewData := make(map[string]interface{})
	var toomany bool

	users := data.FetchUsers()
	if len(users) >= 10 {
		viewData["TooMany"] = true
		toomany = true
	}

	if r.Method == http.MethodPost && !toomany {
		email := r.FormValue("email")

		_, err := mail.ParseAddress(email)

		if err != nil {
			viewData["Error"] = err
		} else {
			data.AddUser(email)
			viewData["SuccessMsg"] = "Successfully added user " + email + ". ✓"
		}
	}

	users = data.FetchUsers()
	if len(users) >= 10 {
		viewData["TooMany"] = true
		toomany = true
	}

	viewData["Users"] = users

	model := views.NewViewModel(nil, viewData)

	views.RenderTemplate(w, "example_adduser", model)
}

func exampleDeleteallHandler(w http.ResponseWriter, r *http.Request) {
	viewData := make(map[string]interface{})

	data.DeleteAllUsers()
	viewData["SuccessMsg"] = "Successfully deleted all users. ✓"

	model := views.NewViewModel(nil, viewData)

	views.RenderTemplate(w, "example_adduser", model)
}

func exampleUploadHandler(w http.ResponseWriter, r *http.Request) {
	viewData := make(map[string]interface{})
	viewData["Title"] = "File Upload"

	if r.Method == http.MethodPost {
		r.ParseMultipartForm(10 << 20)
	
		file, _, err := r.FormFile("file")
		if err != nil {
			log.Println("Error Retrieving the File")
			log.Println(err)
			return
		}
	
		defer file.Close()
	
		tempFile, err := ioutil.TempFile("uploads", "upload-*.png")
		if err != nil {
			log.Println(err)
		}
		defer tempFile.Close()
	
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println(err)
		}
	
		tempFile.Write(fileBytes)
		
		viewData["UploadSuccessMsg"] = "Successfully uploaded file."
	}
	
	model := views.NewViewModel(nil, viewData)
	views.RenderTemplate(w, "example_upload", model)
}
