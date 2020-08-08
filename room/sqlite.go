package room

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/rogaps/hotms/model"
)

type SqliteRepository struct {
	Conn *gorm.DB
}

func NewSqliteRepository(conn *gorm.DB) Repository {
	return &SqliteRepository{conn}
}

func (repository *SqliteRepository) FindAvailableRooms(ctx context.Context, checkIn, checkOut model.Date) (rooms []model.Room, err error) {
	err = repository.Conn.Preload("Rates", "start_date <= ? AND end_date >= ?", checkIn, checkIn).
		Where("status = ?", model.RoomStatusAvailable).
		Where("id NOT IN ?",
			repository.Conn.Table("stay_rooms").
				Select("room_id").
				Joins("JOIN rooms ON rooms.id = stay_rooms.room_id").
				Where("stay_rooms.check_in <= ? AND stay_rooms.check_out >= ?", checkIn, checkIn).
				Or("stay_rooms.check_in < ? AND stay_rooms.check_out >= ?", checkOut, checkOut).
				Or("stay_rooms.check_in <= ? AND stay_rooms.check_out >= ?", checkIn, checkIn).SubQuery()).
		Find(&rooms).Error
	return
}

func (repository *SqliteRepository) FindAll(ctx context.Context) (res []model.Room, err error) {
	err = repository.Conn.Preload("Hotel").Find(&res).Error
	return
}

func (repository *SqliteRepository) FindByID(ctx context.Context, id uint) (res model.Room, err error) {
	err = repository.Conn.Preload("Hotel").First(&res, id).Error
	return
}

func (repository *SqliteRepository) Save(ctx context.Context, room model.Room) (res model.Room, err error) {
	err = repository.Conn.Save(&room).Error
	res = room
	return
}

func (repository *SqliteRepository) Delete(ctx context.Context, room model.Room) (res model.Room, err error) {
	err = repository.Conn.Delete(&room).Error
	res = room
	return
}
