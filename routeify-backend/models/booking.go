package models

type Booking struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty"`
	DriverID    string `json:"driver_id" bson:"driver_id"`
	CabID       string `json:"cab_id" bson:"cab_id"`
	UserID      string `json:"user_id" bson:"user_id"`
	DriverName  string `json:"driver_name" bson:"driver_name"`
	BookingDate string `json:"booking_date" bson:"booking_date"`
	TotalPrice  int    `json:"total_price" bson:"total_price"`
	Status      string `json:"status" bson:"status"`
	// Add other booking fields as needed
}
