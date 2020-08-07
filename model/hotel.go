package model

import "github.com/jinzhu/gorm"

type Hotel struct {
	gorm.Model
	Name    string
	Address string
}
