// models/review.go
package models

type Review struct {
	ID        string `json:"id,omitempty" bson:"_id,omitempty"`
	BookingID string `json:"booking_id" bson:"booking_id"`
	UserID    string `json:"user_id" bson:"user_id"`
	Rating    int    `json:"rating" bson:"rating"`
	Comment   string `json:"comment" bson:"comment"`
	// Add other review fields as needed
}
