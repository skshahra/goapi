package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"bitbucket.org/skshahriarahmed/sh_ra/logs"
)

func (H *DatabaseCollections) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Auth middleWare has called ...")
	cookie, err := r.Cookie("Auth1")
    fmt.Println("🚀 ~ file: auth.go ~ line 48 ~ returnhttp.HandlerFunc ~ cookie : ", cookie)

	if err == nil {
		
		fmt.Println("✨Token is valid Welcome home")
		http.Redirect(w,r,"/",http.StatusSeeOther)
		
		
	}



	data:= struct {
		HOST string
	}{
		HOST: os.Getenv("SERVER_IP"),
	}
	t,err:=template.ParseFiles("templates/login2.html")

	logs.ERROR("Error in template.ParseFiles() ",err)

	t.Execute(w, data)
}