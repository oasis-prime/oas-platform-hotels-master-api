package domain

type PublisherVerifyEmail struct {
	Email string
	Token string
}

type PublisherBookingEmail struct {
	Logo            string `json:"logo"`
	BookingID       string `json:"bookingID"`
	HotelName       string `json:"hotelName"`
	Latitude        string `json:"latitude"`
	Longitude       string `json:"longitude"`
	CategoryName    string `json:"categoryName"`
	HotelAddress    string `json:"hotelAddress"`
	PostalCode      string `json:"postalCode"`
	ZoneName        string `json:"zoneName"`
	DestinationName string `json:"destinationName"`
	Days            string `json:"days"`
	CheckIn         string `json:"checkIn"`
	CheckOut        string `json:"checkOut"`
	RoomType        string `json:"roomType"`
	RoomAmount      string `json:"roomAmount"`
	Adults          string `json:"adults"`
	Cost            string `json:"cost"`
	Email           string `json:"email"`
}
