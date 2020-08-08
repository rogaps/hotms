package room

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rogaps/hotms/model"
)

type Repository interface {
	FindAvailableRooms(ctx context.Context, checkIn, checkOut model.Date) (rooms []model.Room, err error)
	FindAll(ctx context.Context) (res []model.Room, err error)
	FindByID(ctx context.Context, id uint) (res model.Room, err error)
	Save(ctx context.Context, room model.Room) (res model.Room, err error)
	Delete(ctx context.Context, room model.Room) (res model.Room, err error)
}

type Service interface {
	FindAvailableRooms(ctx context.Context, checkIn, checkOut model.Date) (rooms []model.Room, err error)
	FindAll(ctx context.Context) (res []model.Room, err error)
	FindByID(ctx context.Context, id uint) (res model.Room, err error)
	Save(ctx context.Context, room model.Room) (res model.Room, err error)
	Delete(ctx context.Context, room model.Room) (res model.Room, err error)
}

type Delivery interface {
	FindAvailableRooms(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type Router interface {
	AddRoute(r *mux.Router)
}
