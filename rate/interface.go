package rate

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rogaps/hotms/model"
)

type Repository interface {
	FindAll(ctx context.Context) (res []model.RoomRate, err error)
	FindByID(ctx context.Context, id uint) (res model.RoomRate, err error)
	Save(ctx context.Context, rate model.RoomRate) (res model.RoomRate, err error)
	Delete(ctx context.Context, rate model.RoomRate) (res model.RoomRate, err error)
}

type Service interface {
	FindAll(ctx context.Context) (res []model.RoomRate, err error)
	FindByID(ctx context.Context, id uint) (res model.RoomRate, err error)
	Save(ctx context.Context, rate model.RoomRate) (res model.RoomRate, err error)
	Delete(ctx context.Context, rate model.RoomRate) (res model.RoomRate, err error)
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
