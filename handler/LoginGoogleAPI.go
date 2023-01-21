package handler

import (
	// "encoding/json"
	// "fmt"
	"net/http"

	"bitbucket.org/skshahriarahmed/sh_ra/utils"
	// 	"os"
	// 	"time"
	// "bitbucket.org/skshahriarahmed/sh_ra/logs"
	// "bitbucket.org/skshahriarahmed/sh_ra/model"
	// "github.com/golang-jwt/jwt/v4"
)

func (H *DatabaseCollections) GoogleLoginAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// if r.Method != "GET" {
	// 	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }

	// Create oauthState cookie
	oauthState := utils.GenerateStateOauthCookie(w)
	/*
		AuthCodeURL receive state that is a token to protect the user
		from CSRF attacks. You must always provide a non-empty string 
		and validate that it matches the the state query parameter 
		on your redirect callback.
	*/
	u := AppConfig.GoogleLoginConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}
