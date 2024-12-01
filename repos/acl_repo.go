package repos

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	// "go.mongodb.org/mongo-driver/mongo"
	"KSI-BE/config"
	"KSI-BE/model"
)

func GetPermissionsByRole(role string) ([]model.ACL, error) {
	collection := config.GetMongoClient().Database("KSI").Collection("acl")
	cursor, err := collection.Find(context.Background(), bson.M{"role": role})
	if err != nil {
		return nil, err
	}
	var acl []model.ACL
	if err := cursor.All(context.Background(), &acl); err != nil {
		return nil, err
	}
	return acl, nil
}
