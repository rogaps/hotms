package model

import "time"

type Reservation struct {
	ID           uint       `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	OrderID      string     `json:"orderId"`
	CustomerName string     `json:"customerName"`
	Total        float64    `json:"total"`
	StayRooms    []StayRoom `json:"stayRooms"`
}
