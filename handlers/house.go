package handlers

import (
	"encoding/json"
	housedto "housy/dto/house"
	dto "housy/dto/result"
	"housy/models"
	"housy/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerhouse struct {
	HouseRepository repositories.HouseRepository
}

func HandlerHouse(HouseRepository repositories.HouseRepository) *handlerhouse {
	return &handlerhouse{HouseRepository}
}

func (h *handlerhouse) FindHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	house, err := h.HouseRepository.FindHouse()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: house}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerhouse) GetHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	house, err := h.HouseRepository.GetHouse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertGetResponse(house)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerhouse) CreateHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(housedto.CreateHouseRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	house := models.House{
		Name:      request.Name,
		CityID:    request.CityID,
		City:      request.City,
		Address:   request.Address,
		Price:     request.Price,
		TypeRent:  request.TypeRent,
		Ameneties: request.Ameneties,
		BedRoom:   request.BedRoom,
		BathRoom:  request.BathRoom,
	}

	data, err := h.HouseRepository.CreateHouse(house)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertHouseResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func convertHouseResponse(u models.House) housedto.HouseResponse {
	return housedto.HouseResponse{
		ID:        u.ID,
		Name:      u.Name,
		CityID:    u.CityID,
		City:      u.City,
		Address:   u.Address,
		Price:     u.Price,
		TypeRent:  u.TypeRent,
		Ameneties: u.Ameneties,
		BedRoom:   u.BedRoom,
		BathRoom:  u.BathRoom,
	}
}

func convertGetResponse(u models.House) housedto.GetHouseResponse {
	return housedto.GetHouseResponse{
		ID:        u.ID,
		Name:      u.Name,
		CityID:    u.CityID,
		City:      u.City,
		Address:   u.Address,
		Price:     u.Price,
		TypeRent:  u.TypeRent,
		Ameneties: u.Ameneties,
		BedRoom:   u.BedRoom,
		BathRoom:  u.BathRoom,
	}
}
