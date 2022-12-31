package songs

import "go.mongodb.org/mongo-driver/bson/primitive"

type SongFormData struct {
	Name  string `form:"name" validate:"required"`
	URL   string `form:"url" validate:"required,url"`
	Album string `form:"album" validate:"required"`
}

type Song struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string             `json:"name,omitempty" validate:"required"`
	URL   string             `json:"url,omitempty" validate:"required,url"`
	Album string             `json:"album,omitempty" validate:"required"`
}
