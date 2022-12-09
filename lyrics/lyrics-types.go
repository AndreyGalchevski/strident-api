package lyrics

type Lyric struct {
	ID    string `json:"id" bson:"_id"`
	Name  string `json:"name,omitempty" validate:"required"`
	Text  string `json:"text,omitempty" validate:"required"`
	Album string `json:"album,omitempty" validate:"required"`
}
