package members

import "go.mongodb.org/mongo-driver/bson/primitive"

type MemberFormData struct {
	Name       string `form:"name" validate:"required"`
	Instrument string `form:"instrument" validate:"required"`
}

type Member struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Name       string             `json:"name,omitempty" validate:"required"`
	Instrument string             `json:"instrument,omitempty" validate:"required"`
	Image      string             `json:"image,omitempty" validate:"omitempty,url"`
}
