package model

import "github.com/google/uuid"

type Portofolio struct {
	ID       string `json:"id" bson:"_id"`
	DesignImage string `json:"design_image" bson:"design_image"`
	DesignTitle string `json:"design_title" bson:"design_title"`
	DesignDescription string `json:"design_description" bson:"design_description"`
	DesignType string `json:"design_type" bson:"design_type"`
}

func (u *Portofolio) GenerateID() {
	u.ID = uuid.NewString() // Generate UUID
}