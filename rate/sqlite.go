package rate

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

func (repository *SqliteRepository) FindAll(ctx context.Context) (res []model.RoomRate, err error) {
	err = repository.Conn.Find(&res).Error
	return
}

func (repository *SqliteRepository) FindByID(ctx context.Context, id uint) (res model.RoomRate, err error) {
	err = repository.Conn.First(&res, id).Error
	return
}

func (repository *SqliteRepository) Save(ctx context.Context, roomRate model.RoomRate) (res model.RoomRate, err error) {
	err = repository.Conn.Save(&roomRate).Error
	res = roomRate
	return
}

func (repository *SqliteRepository) Delete(ctx context.Context, roomRate model.RoomRate) (res model.RoomRate, err error) {
	err = repository.Conn.Delete(&roomRate).Error
	res = roomRate
	return
}
