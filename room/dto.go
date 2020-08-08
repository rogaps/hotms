package room

import "github.com/rogaps/hotms/model"

type RoomRateDTO struct {
	ID        uint       `json:"id,string,omitempty"`
	StartDate model.Date `json:"startDate"`
	EndDate   model.Date `json:"endDate"`
	Rate      float64    `json:"rate"`
}
