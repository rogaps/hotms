package reservation

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rogaps/hotms/model"
)

type Repository interface {
	FindAll(ctx context.Context) (res []model.Reservation, err error)
	FindByID(ctx context.Context, id uint) (res model.Reservation, err error)
	Save(ctx context.Context, reservation model.Reservation) (res model.Reservation, err error)
	Delete(ctx context.Context, reservation model.Reservation) (res model.Reservation, err error)
}

type Service interface {
	FindAll(ctx context.Context) (res []model.Reservation, err error)
	FindByID(ctx context.Context, id uint) (res model.Reservation, err error)
	Save(ctx context.Context, reservation model.Reservation) (res model.Reservation, err error)
	Delete(ctx context.Context, reservation model.Reservation) (res model.Reservation, err error)
}

type Delivery interface {
	FindAll(w http.ResponseWriter, r *http.Request)
	FindByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type Router interface {
	AddRoute(r *mux.Router)
}
