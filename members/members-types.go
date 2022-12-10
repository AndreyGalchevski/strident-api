package members

import "go.mongodb.org/mongo-driver/bson/primitive"

type Member struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Name       string             `json:"name,omitempty" validate:"required"`
	Instrument string             `json:"instrument,omitempty" validate:"required"`
	Image      string             `json:"image,omitempty" validate:"required,url"`
}
