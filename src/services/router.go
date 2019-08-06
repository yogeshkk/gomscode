package services

import (
	"fmt"
	"gomscode/src/kaau"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	staticDir = "/static/"
)

// Test1 is tset
func Test1() {
	fmt.Println("hello test")
}

// StartServer exported
func StartServer(w http.ResponseWriter, r *http.Request) {
	/*	message := r.URL.Path
			message = strings.TrimPrefix(message, "/")

		logger.LogOut("INFO", r.URL.Path)
		http.ServeFile(w, r, r.URL.Path)
		r := mux.NewRouter()
	*/
}

// NewRouter exported
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	//router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
	//	router.HandleFunc("/web",http.FileServer(http.Dir(staticDir))
	//	router.HandleFunc("/", kubernetes.HomeHandler)
	router.HandleFunc("/login", Middleware(kaau.LoginHandler)).Methods("POST")
	router.HandleFunc("/logout", Middleware(kaau.LogoutPageHandler)).Methods("POST")
	router.HandleFunc("/index", Middleware(kaau.HomeHandler))
	router.HandleFunc("/", Middleware(kaau.LoginPageHandler))
	router.HandleFunc("/app", Middleware(kaau.AppHandler))
//	router.PathPrefix("/css").Handler(http.FileServer(http.Dir("/web/css")))
	router.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))


	return router
}
