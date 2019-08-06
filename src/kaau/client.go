package kaau

import (
	"fmt"
	"gomscode/src/logger"
	"gomscode/src/utility"
	"html/template"
	"net/http"
)

// Todo exported
type Todo struct {
	Title string
	Done  bool
}

// TodoPageData exported
type TodoPageData struct {
	Name string
}

// HomeHandler export
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)

	userName := utility.GetUserName(r)
	data := new(TodoPageData)
	data.Name = userName
	logger.LogOut(data.Name)

	if !utility.IsEmpty(userName) {
		//fmt.Fprintf(w, "Hello, %s you logged to the site.. continue broswing \n", userName)
		tmpl := template.Must(template.ParseFiles("web/templetes/index.html"))
		tmpl.Execute(w, data)
	} else {
		http.Redirect(w, r, "/", 302)
	}

}

// AppHandler export
func AppHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you have requested app: %s\n", r.URL.Path)
}

// LoginHandler export
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	pass := r.FormValue("password")
	redirectTarget := "/"
	if !utility.IsEmpty(name) && !utility.IsEmpty(pass) {
		// Database check for user data!
		IsValidUser := utility.UserIsValid(name, pass)
		if IsValidUser {
			utility.SetCookie(name, w)
			redirectTarget = "/index"
		} else {
			redirectTarget = "/"
		}
	}
	http.Redirect(w, r, redirectTarget, 302)
}

// LoginPageHandler exported
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	userName := utility.GetUserName(r)
	if !utility.IsEmpty(userName) {
		redirectTarget := "/index"
		http.Redirect(w, r, redirectTarget, 302)
	} else {
		tmpl := template.Must(template.ParseFiles("web/templetes/login.html"))
		var data string
		tmpl.Execute(w, data)
	}
}

// LogoutPageHandler exported
func LogoutPageHandler(w http.ResponseWriter, r *http.Request) {
	utility.ClearCookie(w)
	redirectTarget := "/"
	http.Redirect(w, r, redirectTarget, 302)
}
