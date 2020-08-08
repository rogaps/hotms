package hotel

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

func (repository *SqliteRepository) FindAll(ctx context.Context) (res []model.Hotel, err error) {
	err = repository.Conn.Find(&res).Error
	return
}

func (repository *SqliteRepository) FindByID(ctx context.Context, id uint) (res model.Hotel, err error) {
	err = repository.Conn.First(&res, id).Error
	return
}

func (repository *SqliteRepository) Save(ctx context.Context, hotel model.Hotel) (res model.Hotel, err error) {
	err = repository.Conn.Save(&hotel).Error
	res = hotel
	return
}

func (repository *SqliteRepository) Delete(ctx context.Context, hotel model.Hotel) (res model.Hotel, err error) {
	err = repository.Conn.Delete(&hotel).Error
	res = hotel
	return
}
