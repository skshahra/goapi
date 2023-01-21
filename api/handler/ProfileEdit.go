package handler

import (
	// "encoding/json"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"bitbucket.org/skshahriarahmed/sh_ra/logs"
	"bitbucket.org/skshahriarahmed/sh_ra/model"
)

func (H *DatabaseCollections) EditProfile(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Edit Profile has called ...")
cookie, err := r.Cookie("Auth1")
fmt.Println("🚀 ~ file: Home.go ~ line 19 ~ func ~ cookie : ", cookie)

// claims,b := extractClaims(cookie.Value)
	// if b {
	
	// }

	parts := strings.Split(cookie.Value, ".")

	dataByte ,err:=base64.RawURLEncoding.DecodeString(parts[1])
    logs.ERROR("🚀 profile edit.go : ", err)
	var tokenData model.Claims
	err = json.Unmarshal(dataByte, &tokenData)
    logs.ERROR("🚀 ~ file: profile edit.go ~ line 40 ~ func ~ err : ", err)
    fmt.Println("🚀profile edit.go tokenData : ", tokenData)
	
	var UserData model.UserData

	H.MySqlDB.Table("user_data").Select("*").Where("email = ?", tokenData.Email).Scan(&UserData)

	// H.MySqlDB.Model(&model.UserData{Email: tokenData.Email}).First(&UserData)
	fmt.Println("🚀 profile edit.go UserData : ", UserData)


	t,err:=template.ParseFiles("templates/editProfile2.html")

	logs.ERROR("Error in template.ParseFiles() ",err)

	t.Execute(w, UserData)
}