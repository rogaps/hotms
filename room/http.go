package room

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

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

func parseTimes(checkInStr, checkOutStr string) (checkIn, checkOut model.Date, err error) {
	format := "2006-01-02"
	checkInTime, err := time.Parse(format, checkInStr)
	if err != nil {
		return
	}
	checkIn = model.Date(checkInTime)
	checkOutTime, err := time.Parse(format, checkOutStr)
	if err != nil {
		return
	}
	checkOut = model.Date(checkOutTime)
	return
}

func (delivery *HttpDelivery) FindAvailableRooms(w http.ResponseWriter, r *http.Request) {
	checkIn, checkOut, err := parseTimes(r.URL.Query().Get("checkIn"), r.URL.Query().Get("checkOut"))
	if err != nil {
		response.WithError(w, http.StatusBadRequest, err.Error())
		return
	}

	rooms, err := delivery.service.FindAvailableRooms(r.Context(), checkIn, checkOut)
	if err != nil {
		response.WithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.WithJSON(w, http.StatusOK, rooms)
}

func (delivery *HttpDelivery) FindAll(w http.ResponseWriter, r *http.Request) {
	rooms, err := delivery.service.FindAll(r.Context())
	if err != nil {
		response.WithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.WithJSON(w, http.StatusOK, rooms)
}

func (delivery *HttpDelivery) FindByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	room, err := delivery.service.FindByID(r.Context(), uint(id))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.WithError(w, http.StatusNotFound, err.Error())
			return
		}
		response.WithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.WithJSON(w, http.StatusOK, room)
}

func (delivery *HttpDelivery) Create(w http.ResponseWriter, r *http.Request) {
	var room model.Room
	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		response.WithError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := delivery.service.Save(r.Context(), room)
	if err != nil {
		response.WithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.WithJSON(w, http.StatusCreated, res)
}

func (delivery *HttpDelivery) Update(w http.ResponseWriter, r *http.Request) {
	var room model.Room
	err := json.NewDecoder(r.Body).Decode(&room)
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

	room.ID = res.ID
	res, err = delivery.service.Save(r.Context(), room)
	if err != nil {
		response.WithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}

func (delivery *HttpDelivery) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	room, err := delivery.service.FindByID(r.Context(), uint(id))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.WithError(w, http.StatusNotFound, err.Error())
			return
		}
		response.WithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := delivery.service.Delete(r.Context(), room)
	if err != nil {
		response.WithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}
