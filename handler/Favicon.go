package handler

import (
	// "html/template"
	"net/http"

	// "bitbucket.org/skshahriarahmed/sh_ra/logs"
)

func (H *DatabaseCollections) Favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../templates/favicon.ico")
}