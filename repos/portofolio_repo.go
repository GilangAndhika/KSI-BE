package repos

import (
	"context"
	"log"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"KSI-BE/config"
	"KSI-BE/model"
)

func CreatePortofolio(portofolio *model.Portofolio) (string, error) {
	// Generate ID
	if portofolio.ID == "" {
		portofolio.GenerateID()
	}

	// Akses koleksi 'portofolio' dalam database
	collection := config.GetMongoClient().Database("ksi").Collection("portofolio")

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
	collection := config.GetMongoClient().Database("ksi").Collection("portofolio")
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

func GetAllPortofolio() ([]model.Portofolio, error) {
	collection := config.GetMongoClient().Database("ksi").Collection("portofolio")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error fetching portofolios:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var portofolios []model.Portofolio
	for cursor.Next(context.Background()) {
		var portofolio model.Portofolio
		err := cursor.Decode(&portofolio)
		if err != nil {
			log.Println("Error decoding portofolio:", err)
			return nil, err
		}
		portofolios = append(portofolios, portofolio)
	}
	return portofolios, nil
}