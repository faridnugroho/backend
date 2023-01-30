package routes

import (
	"housy/handlers"
	"housy/pkg/mysql"
	"housy/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transaction", h.FindTransaction).Methods("GET")
	// r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
	// r.HandleFunc("/user", h.CreateUser).Methods("POST")
	// r.HandleFunc("/user/{id}", h.UpdateUser).Methods("PATCH")
	// r.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE")
}
