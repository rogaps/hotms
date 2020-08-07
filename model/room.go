package model

import "github.com/jinzhu/gorm"

type Room struct {
	gorm.Model
	HotelID uint
	Hotel   Hotel
	Number  int
}
