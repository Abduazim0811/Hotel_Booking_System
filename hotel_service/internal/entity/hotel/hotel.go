package hotel


type Hotel struct {
    HotelID  string `json:"hotelId" bson:"_id"`
    Name     string `json:"name" bson:"name"`
    Location string `json:"location" bson:"location"`
    Rating   int32  `json:"rating" bson:"rating"`
    Address  string `json:"address" bson:"address"`
}

type HotelRequest struct {
    Name     string `json:"name" bson:"name"`
    Location string `json:"location" bson:"location"`
    Rating   int32  `json:"rating" bson:"rating"`
    Address  string `json:"address" bson:"address"`
}

type HotelResponse struct {
    HotelID string `json:"hotelId" bson:"_id"`
}

type Empty struct{}

type ListHotels struct {
    Hotels []Hotel `json:"hotels" bson:"hotels"`
}

type HotelRes struct {
    Message string `json:"message" bson:"message"`
}

type Room struct {
    RoomID         string `json:"roomId" bson:"roomId"`
    RoomType       string `json:"roomType" bson:"roomType"`
    PricePerNight  int32  `json:"pricePerNight" bson:"pricePerNight"`
    Availability   bool   `json:"availability" bson:"availability"`
}

type RoomRequest struct {
    RoomType      string `json:"roomType" bson:"roomType"`
    PricePerNight int32  `json:"pricePerNight" bson:"pricePerNight"`
    Availability  bool   `json:"availability" bson:"availability"`
}

type RoomResponse struct {
    RoomID string `json:"roomId" bson:"roomId"`
}

type ListRooms struct {
    Rooms []Room `json:"rooms" bson:"rooms"`
}

type RoomRes struct {
    Message string `json:"message" bson:"message"`
}
