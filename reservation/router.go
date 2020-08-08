package reservation

import (
	"net/http"

	"github.com/gorilla/mux"
)

type ReservationRouter struct {
	delivery Delivery
}

func NewRouter(delivery Delivery) Router {
	return &ReservationRouter{delivery}
}

func (router *ReservationRouter) AddRoute(r *mux.Router) {
	r.HandleFunc("/reservations", router.delivery.FindAll).Methods(http.MethodGet)
	r.HandleFunc("/reservations/{id:[0-9]+}", router.delivery.FindByID).Methods(http.MethodGet)
	r.HandleFunc("/reservations", router.delivery.Create).Methods(http.MethodPost)
	r.HandleFunc("/reservations/{id:[0-9]+}", router.delivery.Update).Methods(http.MethodPut)
	r.HandleFunc("/reservations/{id:[0-9]+}", router.delivery.Delete).Methods(http.MethodDelete)
}
