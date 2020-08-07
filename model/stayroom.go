package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type StayRoom struct {
	gorm.Model
	ReservationID uint
	Reservation   Reservation
	RoomID        uint
	Room          Room
	CheckIn       time.Time
	CheckOut      time.Time
}
