// models/vehicle.go
package models

type Vehicle struct {
	ID           string `json:"id,omitempty" bson:"_id,omitempty"`
	Registration string `json:"registration" bson:"registration"`
	Type         string `json:"type" bson:"type"`
	// Add other vehicle fields as needed
}
