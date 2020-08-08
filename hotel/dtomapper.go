package hotel

import "github.com/rogaps/hotms/model"

func ToHotel(dto HotelDTO) model.Hotel {
	return model.Hotel{Name: dto.Name, Address: dto.Address}
}

func ToHoteDTO(hotel model.Hotel) HotelDTO {
	return HotelDTO{Name: hotel.Name, Address: hotel.Address}
}
