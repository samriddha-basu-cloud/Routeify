// models/payment.go
package models

type Payment struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty"`
	BookingID   string `json:"booking_id" bson:"booking_id"`
	Amount      int    `json:"amount" bson:"amount"`
	PaymentDate string `json:"payment_date" bson:"payment_date"`
	Status      string `json:"status" bson:"status"`
	// Add other payment fields as needed
}
