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

func GetUserByUsername(username string) (*model.User, error) {
	collection := config.GetMongoClient().Database("KSI").Collection("users")
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
	// Akses koleKSI 'users' dalam database
	collection := config.GetMongoClient().Database("KSI").Collection("users")

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

func GetUserByID(id string) (*model.User, error) {
	collection := config.GetMongoClient().Database("KSI").Collection("users")
	var user model.User

	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Jika tidak ada user dengan ID ini, return nil
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetAllUser() ([]model.User, error) {
	collection := config.GetMongoClient().Database("KSI").Collection("users")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error fetching users:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var users []model.User
	for cursor.Next(context.Background()) {
		var user model.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Println("Error decoding user:", err)
			return nil, err
		}

		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}

	return users, nil
}

func UpdateUser(id string, user *model.User) (string, error) {
	collection := config.GetMongoClient().Database("KSI").Collection("users")

	// Cek apakah user sudah ada
	existingUser, err := GetUserByID(id)
	if err != nil {
		log.Println("Error checking user:", err)
		return "", err
	}
	if existingUser == nil {
		return "", fmt.Errorf("user with ID '%s' does not exist", id)
	}

	// Hapus _id dari struct user untuk mencegah pengubahan _id
	updateData := bson.M{
		"username": user.Username,
		"email":    user.Email,
		"phone":    user.Phone,
		"role":     user.Role,
	}

	// Update user di MongoDB
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updateData}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Error updating user:", err)
		return "", err
	}

	return id, nil
}


func DeleteUser(id string) error {
	collection := config.GetMongoClient().Database("KSI").Collection("users")

	// Cek apakah user sudah ada
	existingUser, err := GetUserByID(id)
	if err != nil {
		log.Println("Error checking user:", err)
		return err
	}
	if existingUser == nil {
		return fmt.Errorf("user with ID '%s' does not exist", id)
	}

	// Hapus user dari MongoDB
	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		log.Println("Error deleting user:", err)
		return err
	}

	return nil
}
