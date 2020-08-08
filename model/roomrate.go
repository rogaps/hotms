package model

import (
	"time"
)

type RoomRate struct {
	ID        uint      `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Rate      float64   `json:"rate"`
	StartDate Date      `sql:"type:datetime" json:"startDate"`
	EndDate   Date      `sql:"type:datetime" json:"endDate"`
	RoomID    uint      `sql:"type:integer REFERENCES rooms(id)" json:"roomId"`
}
