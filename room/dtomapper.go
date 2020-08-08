package room

import (
	"github.com/rogaps/hotms/model"
)

func ToRoomRate(dto RoomRateDTO) model.RoomRate {
	return model.RoomRate{
		StartDate: dto.StartDate,
		EndDate:   dto.EndDate,
		Rate:      dto.Rate,
	}
}
