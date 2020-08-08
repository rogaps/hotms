package hotel

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rogaps/hotms/helper/response"
	"github.com/rogaps/hotms/model"
)

type HttpDelivery struct {
	service Service
}

func NewHttpDelivery(service Service) Delivery {
	return &HttpDelivery{service}
}

func (delivery *HttpDelivery) FindAll(w http.ResponseWriter, r *http.Request) {
	hotels, err := delivery.service.FindAll(r.Context())
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.WithError(w, http.StatusNotFound, err.Error())
			return
		}
		response.WithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.WithJSON(w, http.StatusOK, hotels)
}

func (delivery *HttpDelivery) FindByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	hotel, err := delivery.service.FindByID(r.Context(), uint(id))
	if err != nil {
		response.WithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.WithJSON(w, http.StatusOK, hotel)
}

func (delivery *HttpDelivery) Create(w http.ResponseWriter, r *http.Request) {
	var hotel model.Hotel
	err := json.NewDecoder(r.Body).Decode(&hotel)
	if err != nil {
		response.WithError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := delivery.service.Save(r.Context(), hotel)
	if err != nil {
		response.WithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.WithJSON(w, http.StatusCreated, res)
}

func (delivery *HttpDelivery) Update(w http.ResponseWriter, r *http.Request) {
	var hotel model.Hotel
	err := json.NewDecoder(r.Body).Decode(&hotel)
	if err != nil {
		response.WithError(w, http.StatusBadRequest, err.Error())
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	res, err := delivery.service.FindByID(r.Context(), uint(id))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.WithError(w, http.StatusNotFound, err.Error())
			return
		}
		response.WithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	hotel.ID = res.ID
	res, err = delivery.service.Save(r.Context(), hotel)
	if err != nil {
		response.WithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}

func (delivery *HttpDelivery) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	hotel, err := delivery.service.FindByID(r.Context(), uint(id))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.WithError(w, http.StatusNotFound, err.Error())
			return
		}
		response.WithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := delivery.service.Delete(r.Context(), hotel)
	if err != nil {
		response.WithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}
