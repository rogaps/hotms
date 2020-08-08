package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rogaps/hotms/hotel"
	"github.com/rogaps/hotms/model"
	"github.com/rogaps/hotms/rate"
	"github.com/rogaps/hotms/reservation"
	"github.com/rogaps/hotms/room"
)

func initDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "hotms.db")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.Exec("PRAGMA foreign_keys = ON;")

	db.AutoMigrate(&model.Hotel{})
	db.AutoMigrate(&model.Room{})
	db.AutoMigrate(&model.Reservation{})
	db.AutoMigrate(&model.RoomRate{})
	db.AutoMigrate(&model.Stay{})
	db.AutoMigrate(&model.StayRoom{})

	return db
}

func main() {
	db := initDB()
	defer db.Close()

	r := mux.NewRouter()

	hotelRouter := hotel.ProvideRouter(db)
	hotelRouter.AddRoute(r)

	roomRouter := room.ProvideRouter(db)
	roomRouter.AddRoute(r)

	rateRouter := rate.ProvideRouter(db)
	rateRouter.AddRoute(r)

	reservationRouter := reservation.ProvideRouter(db)
	reservationRouter.AddRoute(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
