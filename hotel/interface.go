package hotel

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rogaps/hotms/model"
)

type Repository interface {
	FindAll(ctx context.Context) (res []model.Hotel, err error)
	FindByID(ctx context.Context, id uint) (res model.Hotel, err error)
	Save(ctx context.Context, hotel model.Hotel) (res model.Hotel, err error)
	Delete(ctx context.Context, hotel model.Hotel) (res model.Hotel, err error)
}

type Service interface {
	FindAll(ctx context.Context) (res []model.Hotel, err error)
	FindByID(ctx context.Context, id uint) (res model.Hotel, err error)
	Save(ctx context.Context, hotel model.Hotel) (res model.Hotel, err error)
	Delete(ctx context.Context, hotel model.Hotel) (res model.Hotel, err error)
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
