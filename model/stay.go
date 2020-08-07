package model

import "github.com/jinzhu/gorm"

type Stay struct {
	gorm.Model
	StayRoomID uint
	StayRoom   StayRoom
	GuestName  string
}
