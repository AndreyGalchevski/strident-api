package lyrics

import "go.mongodb.org/mongo-driver/bson/primitive"

type LyricFormData struct {
	Name  string `form:"name" validate:"required"`
	Text  string `form:"text" validate:"required"`
	Album string `form:"album" validate:"required"`
}

type Lyric struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string             `json:"name,omitempty" validate:"required"`
	Text  string             `json:"text,omitempty" validate:"required"`
	Album string             `json:"album,omitempty" validate:"required"`
}
