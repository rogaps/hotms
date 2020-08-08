//+build wireinject

package rate

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func ProvideRouter(db *gorm.DB) Router {
	panic(wire.Build(
		NewSqliteRepository,
		NewRoomRateService,
		NewHttpDelivery,
		NewRoomRateRouter,
	))
}
