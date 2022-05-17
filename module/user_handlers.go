package module

import (
	"APLIKASI_1/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	pass := user.Password
	DB.Where("email = ?", user.Email).Find(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if err == nil {
		token, err := model.GenerateJWT(user.Email, user.Password)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		var response = model.JsonResponse{Type: true, Message: "Success Login", Data: []model.User{user}, Token: token.Token}
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []model.User
	DB.Find(&users)
	var response = model.JsonResponse{Type: true, Message: "Success Get data", Data: users}
	json.NewEncoder(w).Encode(response)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var user model.User
	DB.First(&user, params["id"])
	var response = model.JsonResponse{Type: true, Message: "Success Get data by ID", Data: []model.User{user}}
	json.NewEncoder(w).Encode(response)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	user.Password = model.HashPassword(user.Password)
	DB.Create(&user)
	var response = model.JsonResponse{Type: true, Message: "Success Create data", Data: []model.User{user}}
	json.NewEncoder(w).Encode(response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user model.User
	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	user.Password = model.HashPassword(user.Password)
	DB.Save(&user)
	var response = model.JsonResponse{Type: true, Message: "Success Update data", Data: []model.User{user}}
	json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user model.User
	DB.Delete(&user, params["id"])
	var response = model.JsonResponse{Type: true, Message: "User is Deleted Successfully"}
	json.NewEncoder(w).Encode(response)
}
