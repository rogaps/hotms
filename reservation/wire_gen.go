// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package reservation

import (
	"github.com/jinzhu/gorm"
)

// Injectors from wire.go:

func ProvideRouter(db *gorm.DB) Router {
	repository := NewSqliteRepository(db)
	service := NewReservationService(repository)
	delivery := NewHttpDelivery(service)
	router := NewRouter(delivery)
	return router
}
