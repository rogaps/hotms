package model

import "time"

type Stay struct {
	ID         uint      `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	StayRoomID uint      `sql:"type:integer REFERENCES stay_rooms(id)" json:"stayRoomId"`
	GuestName  string    `json:"guestName"`
}
