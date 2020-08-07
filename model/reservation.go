package model

import (
	"github.com/jinzhu/gorm"
)

type Reservation struct {
	gorm.Model
	OrderID      string
	CustomerName string
	Total        float64
}
