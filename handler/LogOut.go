package handler

import (
	// "html/template"
	"net/http"
	// "time"

	// "bitbucket.org/skshahriarahmed/sh_ra/logs"
)

func (H *DatabaseCollections) Logout(w http.ResponseWriter, r *http.Request) {

	cookie := &http.Cookie{
		Name:   "Auth1",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", 302)

	// http.Redirect(w,r,"/login",http.StatusSeeOther)

}