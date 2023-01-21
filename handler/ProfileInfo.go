package handler

import (
	// "encoding/json"
	"html/template"
	"net/http"

	"bitbucket.org/skshahriarahmed/sh_ra/logs"
)

func (H *DatabaseCollections) ProfileInfo(w http.ResponseWriter, r *http.Request) {
	t,err:=template.ParseFiles("templates/userProfile.html")

	logs.ERROR("Error in template.ParseFiles() ",err)

	t.Execute(w, nil)
}