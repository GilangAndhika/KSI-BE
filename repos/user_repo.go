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

func GetUserByUsername(username string) (*model.User, error) {
	collection := config.GetMongoClient().Database("ksi").Collection("users")
	var user model.User
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Jika tidak ada pengguna dengan username ini, return nil
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *model.User) (string, error) {
	// Akses koleksi 'users' dalam database
	collection := config.GetMongoClient().Database("ksi").Collection("users")

	// Cek apakah username sudah ada
	existingUser, err := GetUserByUsername(user.Username)
	if err != nil {
		log.Println("Error checking username:", err)
		return "", err
	}
	if existingUser != nil {
		return "", fmt.Errorf("username '%s' already exists", user.Username)
	}

	// Generate ID untuk user baru jika menggunakan UUID
	user.GenerateID()

	// Insert user baru ke dalam MongoDB
	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Println("Error inserting user:", err)
		return "", err
	}

	// Mengembalikan ID user yang baru disimpan
	id := result.InsertedID
	return id.(string), nil
}
