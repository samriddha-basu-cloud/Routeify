package repository

import (
	"context"
	"log"

	"github.com/samriddha-basu-cloud/Routeify/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db         *mongo.Database
	userCol    *mongo.Collection
	brandCol   *mongo.Collection
	cabCol     *mongo.Collection
	bookingCol *mongo.Collection
	// Define other collections here (e.g., brandCol, cabCol, bookingCol)
)

func ConnectMongoDB(connectionString string) {
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database("Routeify")

	// Initialize collection references
	userCol = db.Collection("users")
	// Initialize collection references
	brandCol = db.Collection("brands")
	cabCol = db.Collection("cabs")
	bookingCol = db.Collection("bookings")
	// Initialize other collections here (e.g., brandCol, cabCol, bookingCol)
}

func CreateUser(user models.User) error {
	_, err := userCol.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func GetUsers() ([]models.User, error) {
	var users []models.User

	cursor, err := userCol.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func UpdateUser(userID string, user models.User) error {
	_, err := userCol.UpdateOne(context.Background(), bson.M{"_id": userID}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(userID string) error {
	_, err := userCol.DeleteOne(context.Background(), bson.M{"_id": userID})
	if err != nil {
		return err
	}
	return nil
}

// CreateBrand creates a new brand
func CreateBrand(brand models.Brand) error {
	_, err := brandCol.InsertOne(context.Background(), brand)
	if err != nil {
		return err
	}
	return nil
}

// GetBrands returns all brands
func GetBrands() ([]models.Brand, error) {
	var brands []models.Brand

	cursor, err := brandCol.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &brands)
	if err != nil {
		return nil, err
	}

	return brands, nil
}

// Define other repository functions similarly for other collections (e.g., brandCol, cabCol, bookingCol)
