package reservation

import (
	"context"

	"github.com/rogaps/hotms/model"
)

type ReservationService struct {
	repository Repository
}

func NewReservationService(repository Repository) Service {
	return &ReservationService{repository}
}

func (service *ReservationService) FindAll(ctx context.Context) (res []model.Reservation, err error) {
	return service.repository.FindAll(ctx)
}

func (service *ReservationService) FindByID(ctx context.Context, id uint) (res model.Reservation, err error) {
	return service.repository.FindByID(ctx, id)
}

func (service *ReservationService) Save(ctx context.Context, reservation model.Reservation) (res model.Reservation, err error) {
	return service.repository.Save(ctx, reservation)
}

func (service *ReservationService) Delete(ctx context.Context, reservation model.Reservation) (res model.Reservation, err error) {
	return service.repository.Delete(ctx, reservation)
}
