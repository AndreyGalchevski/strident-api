package videos

import "go.mongodb.org/mongo-driver/bson/primitive"

type Video struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name,omitempty" validate:"required"`
	URL  string             `json:"url,omitempty" validate:"required,url"`
}
