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

// var path_file = "http://localhost:5000/uploads/"

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

	// for i, p := range house {
	// 	house[i].Image = path_file + p.Image
	// }

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

	// house.Image = path_file + house.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: house}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerhouse) CreateHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// userId := int(userInfo["id"].(float64))

	// dataContex := r.Context().Value("dataFile")
	// filename := dataContex.(string)

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
		// Image:     filename,
		// UserID: userId,
	}

	data, err := h.HouseRepository.CreateHouse(house)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, _ = h.HouseRepository.GetHouse(data.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerhouse) UpdateHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(housedto.UpdateHouseRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	house, err := h.HouseRepository.GetHouse(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Name != "" {
		house.Name = request.Name
	}

	// if request.CityID != "" {
	// 	house.CityID = request.CityID
	// }

	if request.Address != "" {
		house.Address = request.Address
	}

	if request.Price != "" {
		house.Price = request.Price
	}

	if request.TypeRent != "" {
		house.TypeRent = request.TypeRent
	}

	if request.Ameneties != "" {
		house.Ameneties = request.Ameneties
	}

	if request.BedRoom != "" {
		house.BedRoom = request.BedRoom
	}

	if request.BathRoom != "" {
		house.BathRoom = request.BathRoom
	}

	data, err := h.HouseRepository.UpdateHouse(house)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerhouse) DeleteHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	house, err := h.HouseRepository.GetHouse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.HouseRepository.DeleteHouse(house)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
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
