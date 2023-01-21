package route

import (
	// "fmt"
	"fmt"
	"net/http"

	// "net/http"

	"bitbucket.org/skshahriarahmed/sh_ra/database"
	"bitbucket.org/skshahriarahmed/sh_ra/handler"
	"bitbucket.org/skshahriarahmed/sh_ra/middleware"
	"bitbucket.org/skshahriarahmed/sh_ra/model"
	"github.com/gorilla/mux"
)

var H handler.DatabaseCollections


func Router(r *mux.Router) {
    // fmt.Println("🚀 ~ file: route.go ~ line 8 ~ funcRouter ~ (r : ", r)
	Mysqldb:= database.MysqlDBConnection()

	H= database.DatabaseInitialization(Mysqldb)

	H.MySqlDB.AutoMigrate(&model.UserData{})

    fmt.Println("🚀 ~ file: route.go ~ line 14 ~ funcRouter ~ H : ", H)
	handler.LoadConfig()

	r.HandleFunc("/login", H.Login).Methods("GET")
	r.HandleFunc("/api/login", H.LoginAPI).Methods("POST")

	r.HandleFunc("/api/googlelogin", H.GoogleLoginAPI)
	r.HandleFunc("/api/googlelogin/callback", H.GoogleLoginAPICallback)

	r.HandleFunc("/logout", H.Logout).Methods("GET")
	
	r.HandleFunc("/signup", H.Signup).Methods("GET")
	r.HandleFunc("/api/signup", H.SignupAPI).Methods("POST")
	
	// r.Use(middleware.AuthMiddleware)
	
	r.HandleFunc("/", middleware.AuthMiddleware(H.Home)).Methods("GET")
	// r.HandleFunc("/profileinfo", middleware.AuthMiddleware(H.ProfileInfo)).Methods("GET")

	r.HandleFunc("/editprofile", middleware.AuthMiddleware(H.EditProfile)).Methods("GET")
	r.HandleFunc("/api/editprofile", middleware.AuthMiddleware(H.EditProfileAPI)).Methods("POST")
	// r.HandleFunc("/api/editprofile", middleware.AuthMiddleware(H.EditProfileAPI)).Methods("POST")
	r.HandleFunc("/favicon.ico", faviconHandler).Methods("GET")

	// r.PathPrefix("/favicon.ico").Handler(http.FileServer(http.Dir("templates/flag.svg")))
	// r.HandleFunc("/signup", H.Signup).Methods("POST")
	// r.HandleFunc("/signup", H.Signup).Methods("POST")
	
	// r.HandleFunc("/signup", H.Signup).Methods("POST")


}

func serveFiles(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path)
    p := "." + r.URL.Path
    if p == "./" {
        p = "templates/favicon.ico"
    }
    http.ServeFile(w, r, p)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/favicon.ico")
}