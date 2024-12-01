package model

import "github.com/google/uuid"

type User struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Email	 string `json:"email" bson:"email"`
	Phone	 string `json:"phone" bson:"phone"`
	Password string `json:"password" bson:"password"`
	Role     int `json:"role" bson:"role"` // 0 = user, 1 = admin, 2 = designer
}

func (u *User) GenerateID() {
	u.ID = uuid.NewString() // Generate UUID
}