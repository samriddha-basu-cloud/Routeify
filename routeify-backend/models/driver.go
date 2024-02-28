// models/driver.go
package models

type Driver struct {
	ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string `json:"name" bson:"name"`
	License string `json:"license" bson:"license"`
	// Add other driver fields as needed
}
