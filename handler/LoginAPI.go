package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"bitbucket.org/skshahriarahmed/sh_ra/logs"
	"bitbucket.org/skshahriarahmed/sh_ra/model"
	"github.com/golang-jwt/jwt/v4"
)

func (H *DatabaseCollections) LoginAPI(w http.ResponseWriter, r *http.Request) {
	var ReqData model.UserData

	fmt.Println("🚀 ~ file: login.go ~ line 44 ~ func ~ user : ", ReqData)	
	err := json.NewDecoder(r.Body).Decode(&ReqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ReqData.Email == "" || ReqData.Password == "" {
		w.WriteHeader(http.StatusNotFound)
		// json.NewEncoder(w).Encode(`{"status": "Email or Password is empty"}`)
		return
	}
	// result := map[string]interface{}{}
	var result model.UserData
	// H.MySqlDB.Model(&model.UserData{Email: ReqData.Email}).First(&result)
	H.MySqlDB.Table("user_data").Select("*").Where("email = ?", ReqData.Email).Scan(&result)
	fmt.Println("🚀🚀🚀🚀 loginAPI.go  ~ result : ", result)
	if result.Email == ReqData.Email {
		// json.NewEncoder(w).Encode(`{"status": "Email already exist"}`)
		expirationTime := time.Now().Add(time.Hour * 1000)
		myClaim := &model.Claims{
			UserName:   result.Firstname+" "+result.Lastname,
			Email:      result.Email,
			GoogleAuth: false,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		// LOGIN SUCCESSFUL
		// token,err := jwt.ParseWithClaims(jwt.SigningMethodHS256,myClaim)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)

        fmt.Println("🚀 ~ file: login.go ~ line 51 ~ func ~ token : ", token)
		tokenString, err := token.SignedString( []byte(os.Getenv("JWT_SECRET")))
		logs.ERROR("Error in token.SignedString() ", err)
		http.SetCookie(w, &http.Cookie{
			Name:     "Auth1",
			Value:    tokenString,
			Expires:  expirationTime,
			Path:  "/", // <-- this is the important part
			HttpOnly: true,

			// SameSite: SameSiteLaxMode,
		})
		
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(`{"status": "login successfull"}`)
		return
	} else {
		w.WriteHeader(http.StatusNotFound)
		return

	}

	// http.Redirect(w, r, "/", http.StatusSeeOther)
}
