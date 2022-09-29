package domain

type PublisherVerifyEmail struct {
	Email string
	Token string
}

type PublisherBookingEmail struct {
	Adults          string `json:"adults"`
	BookingID       string `json:"bookingID"`
	CategoryName    string `json:"categoryName"`
	CheckIn         string `json:"checkIn"`
	CheckOut        string `json:"checkOut"`
	Cost            string `json:"cost"`
	Days            string `json:"days"`
	DestinationName string `json:"destinationName"`
	Email           string `json:"email"`
	HotelAddress    string `json:"hotelAddress"`
	HotelImage      string `json:"hotelImage"`
	HotelName       string `json:"hotelName"`
	Latitude        string `json:"latitude"`
	Logo            string `json:"logo"`
	Longitude       string `json:"longitude"`
	PostalCode      string `json:"postalCode"`
	RoomAmount      string `json:"roomAmount"`
	RoomType        string `json:"roomType"`
	ZoneName        string `json:"zoneName"`
}
