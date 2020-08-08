package model

import "time"

//type RoomStatus int64
//
//const (
//	RoomStatusAvailable RoomStatus = iota
//	RoomStatusOutOfService
//)
//
//var roomStatuses = [...]string{
//	"available",
//	"out of service",
//}
//
//func (roomStatus RoomStatus) String() string {
//	return roomStatuses[roomStatus]
//}

const (
	RoomStatusAvailable    string = "available"
	RoomStatusOutOfService string = "out of service"
)

type Room struct {
	ID        uint       `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	HotelID   uint       `sql:"type:integer REFERENCES hotels(id)" json:"hotelId"`
	Hotel     Hotel      `json:"hotel"`
	Number    string     `json:"number"`
	Status    string     `json:"status"`
	Rates     []RoomRate `json:"rates"`
}
