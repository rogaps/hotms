//+build wireinject

package room

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func ProvideRouter(db *gorm.DB) Router {
	panic(wire.Build(
		NewSqliteRepository,
		NewRoomService,
		NewHttpDelivery,
		NewRouter,
	))
}
