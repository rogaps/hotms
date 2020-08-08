package model

import (
	"time"
)

type StayRoom struct {
	ID            uint      `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	ReservationID uint      `sql:"type:integer REFERENCES reservations(id)" json:"reservationId"`
	RoomID        uint      `sql:"type:integer REFERENCES rooms(id)" json:"roomId"`
	Room          Room      `gorm:"association_autoupdate:false:association_autocreate:false" json:"room"`
	CheckIn       Date      `sql:"type:datetime" json:"checkIn"`
	CheckOut      Date      `sql:"type:datetime" json:"checkOut"`
	Stays         []Stay    `json:"stays"`
}
