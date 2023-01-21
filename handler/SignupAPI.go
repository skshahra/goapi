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

func (H *DatabaseCollections) SignupAPI(w http.ResponseWriter, r *http.Request) {
	var ReqData model.UserData



	err := json.NewDecoder(r.Body).Decode(&ReqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("🚀 ~ file: login.go ~ line 44 ~ func ~ user : ", ReqData)

	// result := map[string]interface{}{}
	var result model.UserData
	// H.MySqlDB.Model(&model.UserData{Email: ReqData.Email}).First(&result)
	H.MySqlDB.Table("user_data").Select("*").Where("email = ?", ReqData.Email).Scan(&result)
	fmt.Println("🚀 ~ file: SignupAPI.go ~ line 36 ~ func ~ result : ", result)
	if result.Email == ReqData.Email {
		w.WriteHeader(http.StatusNotFound)
		// json.NewEncoder(w).Encode(`{"status": "Email already exist"}`)
		return
	} else {
		res := H.MySqlDB.Create(&ReqData)
		fmt.Println("🚀 ~ file: SignupAPI.go ~ line 39 ~ func ~ result : ", res)

		expirationTime := time.Now().Add(time.Hour * 1000)
		myClaim := &model.Claims{
			UserName:   ReqData.Firstname+" "+ReqData.Lastname,
			Email:      ReqData.Email,
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
		json.NewEncoder(w).Encode(`{"status": "Registration successfull"}`)
	}

}
