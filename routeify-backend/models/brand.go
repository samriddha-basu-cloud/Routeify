package models

type Brand struct {
	ID   string `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	// Add other brand fields as needed
}
