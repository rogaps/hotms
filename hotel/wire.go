//+build wireinject

package hotel

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func ProvideRouter(db *gorm.DB) Router {
	panic(wire.Build(
		NewSqliteRepository,
		NewHotelService,
		NewHttpDelivery,
		NewRouter,
	))
}
