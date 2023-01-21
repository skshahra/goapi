package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/skshahriarahmed/sh_ra/model"
)

func (H *DatabaseCollections) EditProfileAPI(w http.ResponseWriter, r *http.Request) {
	var ReqData model.UserData



	err := json.NewDecoder(r.Body).Decode(&ReqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("🚀 ~ file: ProfileEditAPI.go ~ line 17 ~ func ~ ReqData : ", ReqData)

	H.MySqlDB.Model(model.UserData{}).
	Where("email = ?", ReqData.Email).
	Updates(ReqData)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(`{status: "Edit Profile api"}`)
}