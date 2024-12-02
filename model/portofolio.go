package model

import "github.com/google/uuid"

type Portofolio struct {
	ID       string `json:"id" bson:"_id"`
	UserID string `json:"user_id" bson:"user_id"`
	Username string `json:"username" bson:"username"`
	Email string `json:"email" bson:"email"`
	Phone string `json:"phone" bson:"phone"`
	Role int `json:"role" bson:"role"`
	DesignImage string `json:"design_image" bson:"design_image"`
	DesignTitle string `json:"design_title" bson:"design_title"`
	DesignDescription string `json:"design_description" bson:"design_description"`
	DesignType string `json:"design_type" bson:"design_type"`
}

func (u *Portofolio) GenerateID() {
	u.ID = uuid.NewString() // Generate UUID
}

// FillUserDetails populates order fields with data from a User
func (u *Portofolio) FillUserDetails(user *User) {
	u.UserID = user.ID
	u.Username = user.Username
	u.Email = user.Email
	u.Phone = user.Phone
	u.Role = user.Role
}