package repos

import (
	"KSI-BE/config"
	"KSI-BE/model"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePortofolio(portofolio *model.Portofolio) (string, error) {
	// Generate ID
	if portofolio.ID == "" {
		portofolio.GenerateID()
	}

	// Akses koleKSI 'portofolio' dalam database
	collection := config.GetMongoClient().Database("KSI").Collection("portofolio")

	// Cek apakah portofolio sudah ada
	existingPortofolio, err := GetPortofolioByID(portofolio.ID)
	if err != nil {
		log.Println("Error checking portofolio:", err)
		return "", err
	}
	if existingPortofolio != nil {
		return "", fmt.Errorf("portofolio '%s' already exists", portofolio.ID)
	}

	// Insert portofolio baru ke dalam MongoDB
	result, err := collection.InsertOne(context.Background(), portofolio)
	if err != nil {
		log.Println("Error inserting portofolio:", err)
		return "", err
	}

	// Mengembalikan ID portofolio yang baru disimpan
	id := result.InsertedID.(string)
	return id, nil
}

func GetPortofolioByID(id string) (*model.Portofolio, error) {
	collection := config.GetMongoClient().Database("KSI").Collection("portofolio")
	var portofolio model.Portofolio

	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&portofolio)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Jika tidak ada portofolio dengan ID ini, return nil
			return nil, nil
		}
		return nil, err
	}
	return &portofolio, nil
}

func GetAllPortofolio() ([]map[string]interface{}, error) {
	collection := config.GetMongoClient().Database("KSI").Collection("portofolio")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error fetching portofolios:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var portofolios []map[string]interface{}
	for cursor.Next(context.Background()) {
		var portofolio model.Portofolio
		err := cursor.Decode(&portofolio)
		if err != nil {
			log.Println("Error decoding portofolio:", err)
			return nil, err
		}

		// Exclude the Password field and construct a map
		portofolioMap := map[string]interface{}{
			"id":          portofolio.ID,
			"title":       portofolio.DesignTitle,
			"description": portofolio.DesignDescription,
			"image":       portofolio.DesignImage,
			"type":        portofolio.DesignType,
		}

		portofolios = append(portofolios, portofolioMap)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}

	return portofolios, nil
}
