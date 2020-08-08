package room

import (
	"net/http"

	"github.com/gorilla/mux"
)

type RoomRouter struct {
	delivery Delivery
}

func NewRouter(delivery Delivery) Router {
	return &RoomRouter{delivery}
}

func (router *RoomRouter) AddRoute(r *mux.Router) {
	r.HandleFunc("/search", router.delivery.FindAvailableRooms).Methods(http.MethodGet)
	r.HandleFunc("/rooms", router.delivery.FindAll).Methods(http.MethodGet)
	r.HandleFunc("/rooms/{id:[0-9]+}", router.delivery.FindByID).Methods(http.MethodGet)
	r.HandleFunc("/rooms", router.delivery.Create).Methods(http.MethodPost)
	r.HandleFunc("/rooms/{id:[0-9]+}", router.delivery.Update).Methods(http.MethodPut)
	r.HandleFunc("/rooms/{id:[0-9]+}", router.delivery.Delete).Methods(http.MethodDelete)
}
