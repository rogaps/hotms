package rate

import (
	"net/http"

	"github.com/gorilla/mux"
)

type RoomRateRouter struct {
	delivery Delivery
}

func NewRoomRateRouter(delivery Delivery) Router {
	return &RoomRateRouter{delivery}
}

func (router *RoomRateRouter) AddRoute(r *mux.Router) {
	r.HandleFunc("/rates", router.delivery.FindAll).Methods(http.MethodGet)
	r.HandleFunc("/rates/{id:[0-9]+}", router.delivery.FindByID).Methods(http.MethodGet)
	r.HandleFunc("/rates", router.delivery.Create).Methods(http.MethodPost)
	r.HandleFunc("/rates/{id:[0-9]+}", router.delivery.Update).Methods(http.MethodPut)
	r.HandleFunc("/rates/{id:[0-9]+}", router.delivery.Delete).Methods(http.MethodDelete)
}
