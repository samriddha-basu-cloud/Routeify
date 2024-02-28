// models/location.go
package models

type Location struct {
	ID        string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string  `json:"name" bson:"name"`
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
	// Add other location fields as needed
}
