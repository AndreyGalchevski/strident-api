package videos

import "go.mongodb.org/mongo-driver/bson/primitive"

type VideoFormData struct {
	Name string `form:"name" validate:"required"`
	URL  string `form:"url" validate:"required,url"`
}

type Video struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name,omitempty" validate:"required"`
	URL  string             `json:"url,omitempty" validate:"required,url"`
}
