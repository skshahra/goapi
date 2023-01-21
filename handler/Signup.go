package handler

import (
	// "encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"bitbucket.org/skshahriarahmed/sh_ra/logs"
	// "bitbucket.org/skshahriarahmed/sh_ra/model"
	// "github.com/golang-jwt/jwt/v4"
)

func (H *DatabaseCollections) Signup(w http.ResponseWriter, r *http.Request) {

	// 	w.Header().Set("Access-Control-Allow-Origin","*")
	// w.Header().Add("Content-Type","application/json")
	// an example API handler
	fmt.Println("Auth middleWare has called ...")
	cookie, err := r.Cookie("Auth1")
    fmt.Println("🚀 ~ file: auth.go ~ line 48 ~ returnhttp.HandlerFunc ~ cookie : ", cookie)

	if err == nil {
		
		fmt.Println("✨Token is valid Welcome home")
		http.Redirect(w,r,"/",http.StatusSeeOther)
		
		
	}





	//////
	data:= struct {
		HOST string
	}{
		HOST: os.Getenv("IP")+":"+os.Getenv("PORT"),
	}

	t,err:=template.ParseFiles("templates/signup2.html")

	logs.ERROR("Error in template.ParseFiles() ",err)

	t.Execute(w, data)
}