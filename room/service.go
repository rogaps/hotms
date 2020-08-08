package room

import (
	"context"

	"github.com/rogaps/hotms/model"
)

type RoomService struct {
	repository Repository
}

func NewRoomService(repository Repository) Service {
	return &RoomService{repository}
}

func (service *RoomService) FindAvailableRooms(ctx context.Context, checkIn, checkOut model.Date) (rooms []model.Room, err error) {
	return service.repository.FindAvailableRooms(ctx, checkIn, checkOut)
}

func (service *RoomService) FindAll(ctx context.Context) (res []model.Room, err error) {
	return service.repository.FindAll(ctx)
}

func (service *RoomService) FindByID(ctx context.Context, id uint) (res model.Room, err error) {
	return service.repository.FindByID(ctx, id)
}

func (service *RoomService) Save(ctx context.Context, room model.Room) (res model.Room, err error) {
	return service.repository.Save(ctx, room)
}

func (service *RoomService) Delete(ctx context.Context, room model.Room) (res model.Room, err error) {
	return service.repository.Delete(ctx, room)
}
