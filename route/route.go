package route

import (
	"APLIKASI_1/model"
	"APLIKASI_1/module"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/login", module.LoginUser).Methods("POST")
	r.HandleFunc("/users", model.JWTMiddleware(module.GetUsers)).Methods("GET")
	r.HandleFunc("/users/{id}", model.JWTMiddleware(module.GetUser)).Methods("GET")
	r.HandleFunc("/users", module.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", model.JWTMiddleware(module.UpdateUser)).Methods("PUT")
	r.HandleFunc("/users/{id}", model.JWTMiddleware(module.DeleteUser)).Methods("DELETE")

	r.HandleFunc("/presensi/{id}", model.JWTMiddleware(module.GetPresen)).Methods("GET")
	r.HandleFunc("/presensi", model.JWTMiddleware(module.Presen)).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", r))

}
