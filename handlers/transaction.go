package handlers

import (
	"encoding/json"
	// housedto "housy/dto/house"
	dto "housy/dto/result"
	// usersdto "housy/dto/users"
	// "housy/models"
	"housy/repositories"
	"net/http"
	// "strconv"
	// "github.com/go-playground/validator/v10"
	// "github.com/gorilla/mux"
)

type handlertransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlertransaction {
	return &handlertransaction{TransactionRepository}
}

func (h *handlertransaction) FindTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transaction, err := h.TransactionRepository.FindTransaction()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transaction}
	json.NewEncoder(w).Encode(response)
}

// func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])

// 	user, err := h.UserRepository.GetUser(id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: user}
// 	json.NewEncoder(w).Encode(response)
// }

// func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	request := new(usersdto.CreateUserRequest)
// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	validation := validator.New()
// 	err := validation.Struct(request)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	user := models.User{
// 		Fullname: request.Fullname,
// 		Username: request.Username,
// 		Email:    request.Email,
// 		Password: request.Password,
// 		ListAsID: request.ListAsID,
// 		ListAs:   request.ListAs,
// 		Gender:   request.Gender,
// 		Address:  request.Address,
// 	}

// 	data, err := h.UserRepository.CreateUser(user)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
// 	json.NewEncoder(w).Encode(response)
// }

// func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	request := new(usersdto.UpdateUserRequest)
// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])
// 	user, err := h.UserRepository.GetUser(int(id))
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	if request.Fullname != "" {
// 		user.Fullname = request.Fullname
// 	}

// 	if request.Username != "" {
// 		user.Username = request.Username
// 	}

// 	if request.Email != "" {
// 		user.Email = request.Email
// 	}

// 	if request.Password != "" {
// 		user.Password = request.Password
// 	}

// 	// if request.ListAsID != "" {
// 	// 	user.ListAsID = request.ListAsID
// 	// }

// 	if request.Gender != "" {
// 		user.Gender = request.Gender
// 	}

// 	if request.Address != "" {
// 		user.Address = request.Address
// 	}

// 	data, err := h.UserRepository.UpdateUser(user)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
// 	json.NewEncoder(w).Encode(response)
// }

// func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])
// 	user, err := h.UserRepository.GetUser(id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	data, err := h.UserRepository.DeleteUser(user)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
// 	json.NewEncoder(w).Encode(response)
// }

// func convertResponse(u models.User) usersdto.UserResponse {
// 	return usersdto.UserResponse{
// 		ID:       u.ID,
// 		Fullname: u.Fullname,
// 		Username: u.Username,
// 		Email:    u.Email,
// 		Password: u.Password,
// 		ListAsID: u.ListAsID,
// 		ListAs:   u.ListAs,
// 		Gender:   u.Gender,
// 		Address:  u.Address,
// 	}
// }
