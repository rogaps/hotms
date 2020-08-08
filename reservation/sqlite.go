package reservation

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

func (repository *SqliteRepository) FindAll(ctx context.Context) (res []model.Reservation, err error) {
	err = repository.Conn.Set("gorm:auto_preload", true).Find(&res).Error
	return
}

func (repository *SqliteRepository) FindByID(ctx context.Context, id uint) (res model.Reservation, err error) {
	err = repository.Conn.Set("gorm:auto_preload", true).First(&res, id).Error
	return
}

func (repository *SqliteRepository) Save(ctx context.Context, reservation model.Reservation) (res model.Reservation, err error) {
	err = repository.Conn.Set("gorm:auto_preload", true).Save(&reservation).Error
	res = reservation
	return
}

func (repository *SqliteRepository) Delete(ctx context.Context, reservation model.Reservation) (res model.Reservation, err error) {
	err = repository.Conn.Set("gorm:auto_preload", true).Delete(&reservation).Error
	res = reservation
	return
}
