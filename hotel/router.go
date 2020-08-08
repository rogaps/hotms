package hotel

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HotelRouter struct {
	delivery Delivery
}

func NewRouter(delivery Delivery) Router {
	return &HotelRouter{delivery}
}

func (router *HotelRouter) AddRoute(r *mux.Router) {
	r.HandleFunc("/hotels", router.delivery.FindAll).Methods(http.MethodGet)
	r.HandleFunc("/hotels/{id:[0-9]+}", router.delivery.FindByID).Methods(http.MethodGet)
	r.HandleFunc("/hotels", router.delivery.Create).Methods(http.MethodPost)
	r.HandleFunc("/hotels/{id:[0-9]+}", router.delivery.Update).Methods(http.MethodPut)
	r.HandleFunc("/hotels/{id:[0-9]+}", router.delivery.Delete).Methods(http.MethodDelete)
}
