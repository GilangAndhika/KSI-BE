package model

type ACL struct {
	Role     string `json:"role" bson:"role"`
	Resource string `json:"resource" bson:"resource"`
	Read     bool   `json:"read" bson:"read"`
	Write    bool   `json:"write" bson:"write"`
	Delete   bool   `json:"delete" bson:"delete"`
}
