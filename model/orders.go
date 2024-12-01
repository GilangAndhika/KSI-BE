package model

import (
	"time"
	"github.com/google/uuid"
)

type Orders struct {
	ID          	 string    `json:"id" bson:"_id"`
	UserID      	 string    `json:"user_id" bson:"user_id"` // Referensi ke ID pengguna
	Username    	 string    `json:"username" bson:"username"`
	Email       	 string    `json:"email" bson:"email"`
	Phone       	 string    `json:"phone" bson:"phone"`
	DesignOrderType  string    `json:"design_order_type" bson:"design_order_type"`
	Reference   	 string    `json:"reference" bson:"reference"` // Referensi file gambar, URL, dll.
	OrdersDate  	 time.Time `json:"orders_date" bson:"orders_date"`
}

// GenerateID generates a UUID for the Orders ID
func (o *Orders) GenerateID() {
	o.ID = uuid.NewString()
}

// FillUserDetails populates order fields with data from a User
func (o *Orders) FillUserDetails(user *User) {
	o.UserID = user.ID
	o.Username = user.Username
	o.Email = user.Email
	o.Phone = user.Phone
}
