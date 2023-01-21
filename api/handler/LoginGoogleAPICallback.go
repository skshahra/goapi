package handler

import (
	// "encoding/json"
	// "fmt"
	// "context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"bitbucket.org/skshahriarahmed/sh_ra/logs"
	"bitbucket.org/skshahriarahmed/sh_ra/model"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
	// "golang.org/x/oauth2"
	// 	"os"
	// 	"time"
	// "bitbucket.org/skshahriarahmed/sh_ra/logs"
	// "bitbucket.org/skshahriarahmed/sh_ra/model"
	// "github.com/golang-jwt/jwt/v4"
)

// type GoogleAuthData struct {
// 	id             string `json:"id"`
// 	email          string `json:"email"`
// 	verified_email bool   `json:"verified_email"`
// 	name           string `json:"name"`
// 	given_name     string `json:"given_name"`
// 	family_name    string `json:"family_name"`
// 	picture        string `json:"picture"`
// 	locale         string `json:"locale"`
// }

func (H *DatabaseCollections) GoogleLoginAPICallback(w http.ResponseWriter, r *http.Request) {

	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	fmt.Println("🚀 ~ file: LoginGoogleAPICallback.go ~ line 23 ~ func ~ content : ", content)
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	var ReqData map[string]interface{}

	err = json.Unmarshal(content, &ReqData)
	logs.ERROR("🚀 ~ file: LoginGoogleAPICallback.go ~ line 49 ~ func ~ err : ", err)

	fmt.Println("🚀 ~ file: LoginGoogleAPICallback.go ~ line 45 ~ func ~ Data : ", ReqData)
	if ReqData["email"] != "" {

		var result model.UserData
		// H.MySqlDB.Model(&model.UserData{Email: ReqData["email"].(string)}).First(&result)
	H.MySqlDB.Table("user_data").Select("*").Where("email = ?", ReqData["email"].(string)).Scan(&result)

		fmt.Println("🚀 ~ file: loginAPI.go ~ line 36 ~ func ~ result : ", result)
		if result.Email == ReqData["email"].(string) && result.GoogleAuth == true {
			expirationTime := time.Now().Add(time.Hour * 1000)
			myClaim := &model.Claims{
				UserName:   result.Firstname + " " + result.Lastname,
				Email:      result.Email,
				GoogleAuth: true,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}
			// LOGIN SUCCESSFUL
			// token,err := jwt.ParseWithClaims(jwt.SigningMethodHS256,myClaim)
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)

			fmt.Println("🚀 ~ file: login.go ~ line 51 ~ func ~ token : ", token)
			tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
			logs.ERROR("Error in token.SignedString() ", err)
			http.SetCookie(w, &http.Cookie{
				Name:     "Auth1",
				Value:    tokenString,
				Expires:  expirationTime,
				Path:     "/", // <-- this is the important part
				HttpOnly: true,
			})

			// w.WriteHeader(http.StatusOK)
			// json.NewEncoder(w).Encode(`{"status": "login successfull"}`)
			// return
		} else {
			var result model.UserData
			result.Firstname = ReqData["given_name"].(string)
			result.Lastname = ReqData["family_name"].(string)
			result.Email = ReqData["email"].(string)
			result.Password = ""
			result.GoogleAuth = true

			res := H.MySqlDB.Create(&result)
			fmt.Println("🚀 ~ file: LoginGoogleAPICallback.go ~ line 93 ~ func ~ res : ", res)

			expirationTime := time.Now().Add(time.Hour * 1000)
			myClaim := &model.Claims{
				UserName:   ReqData["given_name"].(string) + " " + ReqData["family_name"].(string),
				Email:      ReqData["email"].(string),
				GoogleAuth: true,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}
			// LOGIN SUCCESSFUL
			// token,err := jwt.ParseWithClaims(jwt.SigningMethodHS256,myClaim)
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)

			fmt.Println("🚀 ~ file: login.go ~ line 51 ~ func ~ token : ", token)
			tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
			logs.ERROR("Error in token.SignedString() ", err)
			http.SetCookie(w, &http.Cookie{
				Name:     "Auth1",
				Value:    tokenString,
				Expires:  expirationTime,
				Path:     "/", // <-- this is the important part
				HttpOnly: true,

				// SameSite: SameSiteLaxMode,
			})
			// w.WriteHeader(http.StatusOK)
			// json.NewEncoder(w).Encode(`{"status": "Registration successfull"}`)
			// w.WriteHeader(http.StatusNotFound)
			// return

		}
	}
	// fmt.Fprintf(w, "Content: %s\n", content)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func getUserInfo(state string, code string) ([]byte, error) {
	// if state != "" {
	// 	return nil, fmt.Errorf("invalid oauth state")
	// }

	token, err := AppConfig.GoogleLoginConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	fmt.Println("🚀 ~ file: LoginGoogleAPICallback.go ~ line 57 ~ funcgetUserInfo ~ contents : ", contents)
	return contents, nil
}

// func (H *DatabaseCollections) GoogleLoginAPICallback(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	// if r.Method != "GET" {
// 	// 	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 	// 	return
// 	// }

// 	// get oauth state from cookie for this user
// 	oauthState, _ := r.Cookie("oauthstate")
// 	state := r.FormValue("state")
// 	code := r.FormValue("code")
// 	w.Header().Add("content-type", "application/json")

// 	// ERROR : Invalid OAuth State
// 	if state != oauthState.Value {
// 		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		fmt.Fprintf(w, "invalid oauth google state")
// 		return
// 	}

// 	// Exchange Auth Code for Tokens
// 	token, err := AppConfig.GoogleLoginConfig.Exchange(
// 		context.Background(), code)

// 	// ERROR : Auth Code Exchange Failed
// 	if err != nil {
// 		fmt.Fprintf(w, "falied code exchange: %s", err.Error())
// 		return
// 	}

// 	// Fetch User Data from google server
// 	response, err := http.Get(OauthGoogleUrlAPI + token.AccessToken)

// 	// ERROR : Unable to get user data from google
// 	if err != nil {
// 		fmt.Fprintf(w, "failed getting user info: %s", err.Error())
// 		return
// 	}

// 	// Parse user data JSON Object
// 	defer response.Body.Close()
// 	contents, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "failed read response: %s", err.Error())
// 		return
// 	}

// 	// send back response to browser
// 	fmt.Fprintln(w, string(contents))

// }
