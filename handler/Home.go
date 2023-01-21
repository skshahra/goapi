package handler

import (
	// "encoding/json"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	// "log"
	"net/http"
	// "os"
	"strings"

	// "github.com/golang-jwt/jwt/v4"

	"bitbucket.org/skshahriarahmed/sh_ra/logs"
	"bitbucket.org/skshahriarahmed/sh_ra/model"
)

func (H *DatabaseCollections) Home(w http.ResponseWriter, r *http.Request) {
// 
// tmp,_:=template.ParseFiles("templates/index.html")

// tmp.Execute(w, nil)
fmt.Println("Auth middleWare has called ...")
cookie, err := r.Cookie("Auth1")
fmt.Println("🚀 ~ file: Home.go ~ line 19 ~ func ~ cookie : ", cookie)

// claims,b := extractClaims(cookie.Value)
	// if b {
	
	// }

	parts := strings.Split(cookie.Value, ".")

	dataByte ,err:=base64.RawURLEncoding.DecodeString(parts[1])
    logs.ERROR("🚀 ~ file: Home.go ~ line 36 ~ func ~ err : ", err)
	var tokenData model.Claims
	err = json.Unmarshal(dataByte, &tokenData)
    logs.ERROR("🚀 ~ file: Home.go ~ line 40 ~ func ~ err : ", err)
    fmt.Println("🚀 ~ file: Home.go ~ line 51 ~ func ~ tokenData : ", tokenData)
	
	var UserData model.UserData

	// H.MySqlDB.Model(&model.UserData{Email: tokenData.Email}).First(&UserData)

	H.MySqlDB.Table("user_data").Select("*").Where("email = ?", tokenData.Email).Scan(&UserData)
	fmt.Println("🚀🚀🚀🚀 file: Home.go  UserData : ", UserData)



	fileName := "templates/index.html"
	t, err := template.ParseFiles(fileName)
	logs.ERROR("Error in template.ParseFiles() ", err)
	t.Execute(w, UserData)
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(`{status: "login"}`)
}


// func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
// 	hmacSecretString := os.Getenv("JWT_SECRET")
// 	hmacSecret := []byte(hmacSecretString)
// 	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
// 		 // check token signing method etc
// 		 return hmacSecret, nil
// 	})

// 	if err != nil {
// 		return nil, false
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		return claims, true
// 	} else {
// 		log.Printf("Invalid JWT Token")
// 		return nil, false
// 	}
// }