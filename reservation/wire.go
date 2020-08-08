//+build wireinject

package reservation

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func ProvideRouter(db *gorm.DB) Router {
	panic(wire.Build(
		NewSqliteRepository,
		NewReservationService,
		NewHttpDelivery,
		NewRouter,
	))
}
