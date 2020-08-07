package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type RoomRate struct {
	gorm.Model
	Rate      float64
	StartDate time.Time
	EndDate   time.Time
	RoomID    uint
	Room      Room
}
