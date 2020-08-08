package rate

import (
	"context"

	"github.com/rogaps/hotms/model"
)

type RoomRateService struct {
	repository Repository
}

func NewRoomRateService(repository Repository) Service {
	return &RoomRateService{repository}
}

func (service *RoomRateService) FindAll(ctx context.Context) (res []model.RoomRate, err error) {
	return service.repository.FindAll(ctx)
}

func (service *RoomRateService) FindByID(ctx context.Context, id uint) (res model.RoomRate, err error) {
	return service.repository.FindByID(ctx, id)
}

func (service *RoomRateService) Save(ctx context.Context, roomRate model.RoomRate) (res model.RoomRate, err error) {
	return service.repository.Save(ctx, roomRate)
}

func (service *RoomRateService) Delete(ctx context.Context, roomRate model.RoomRate) (res model.RoomRate, err error) {
	return service.repository.Delete(ctx, roomRate)
}
