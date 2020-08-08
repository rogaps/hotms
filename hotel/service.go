package hotel

import (
	"context"

	"github.com/rogaps/hotms/model"
)

type HotelService struct {
	repository Repository
}

func NewHotelService(repository Repository) Service {
	return &HotelService{repository}
}

func (service *HotelService) FindAll(ctx context.Context) (res []model.Hotel, err error) {
	return service.repository.FindAll(ctx)
}

func (service *HotelService) FindByID(ctx context.Context, id uint) (res model.Hotel, err error) {
	return service.repository.FindByID(ctx, id)
}

func (service *HotelService) Save(ctx context.Context, hotel model.Hotel) (res model.Hotel, err error) {
	return service.repository.Save(ctx, hotel)
}

func (service *HotelService) Delete(ctx context.Context, hotel model.Hotel) (res model.Hotel, err error) {
	return service.repository.Delete(ctx, hotel)
}
