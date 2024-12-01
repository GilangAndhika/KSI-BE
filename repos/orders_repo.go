package repos

import (
	"context"
	"fmt"
	"log"

	"KSI-BE/config"
	"KSI-BE/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateOrder(order *model.Orders) (string, error) {
	// Generate ID
	if order.ID == "" {
		order.GenerateID()
	}

	// Akses koleKSI 'orders' dalam database
	collection := config.GetMongoClient().Database("KSI").Collection("orders")

	// Cek apakah order sudah ada
	existingOrder, err := GetOrderByID(order.ID)
	if err != nil {
		log.Println("Error checking order:", err)
		return "", err
	}
	if existingOrder != nil {
		return "", fmt.Errorf("order '%s' already exists", order.ID)
	}

	// Insert order baru ke dalam MongoDB
	result, err := collection.InsertOne(context.Background(), order)
	if err != nil {
		log.Println("Error inserting order:", err)
		return "", err
	}

	// Mengembalikan ID order yang baru disimpan
	id := result.InsertedID.(string)
	return id, nil
}

func GetOrderByID(id string) (*model.Orders, error) {
	collection := config.GetMongoClient().Database("KSI").Collection("orders")
	var order model.Orders

	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Jika tidak ada order dengan ID ini, return nil
			return nil, nil
		}
		return nil, err
	}
	return &order, nil
}
