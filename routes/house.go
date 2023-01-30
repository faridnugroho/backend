package routes

import (
	"housy/handlers"
	"housy/pkg/mysql"
	"housy/repositories"

	"github.com/gorilla/mux"
)

func HouseRoutes(r *mux.Router) {
	houseRepository := repositories.RepositoryHouse(mysql.DB)
	h := handlers.HandlerHouse(houseRepository)

	r.HandleFunc("/houses", h.FindHouse).Methods("GET")
	r.HandleFunc("/house/{id}", h.GetHouse).Methods("GET")
	r.HandleFunc("/house", h.CreateHouse).Methods("POST")
	// r.HandleFunc("/user/{id}", h.UpdateUser).Methods("PATCH")
	// r.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE")
}
