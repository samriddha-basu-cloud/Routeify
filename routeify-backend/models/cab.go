package models

type Cab struct {
	ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	BrandID string `json:"brand_id" bson:"brand_id"`
	Model   string `json:"model" bson:"model"`
	Year    int    `json:"year" bson:"year"`
	// Add other cab fields as needed
}
