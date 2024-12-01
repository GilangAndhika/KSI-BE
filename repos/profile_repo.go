package repos

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"KSI-BE/config"
	"KSI-BE/model"
)

func GetAllProfile() ([]map[string]interface{}, error) {
	collection := config.GetMongoClient().Database("ksi").Collection("user")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error fetching profiles:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var profiles []map[string]interface{}
	for cursor.Next(context.Background()) {
		var user model.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Println("Error decoding profile:", err)
			return nil, err
		}

		// Exclude the Password field and construct a map
		profile := map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"phone":    user.Phone,
			"role":     user.Role,
		}

		profiles = append(profiles, profile)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}

	return profiles, nil
}