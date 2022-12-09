package songs

type Song struct {
	ID    string `json:"id" bson:"_id"`
	Name  string `json:"name,omitempty" validate:"required"`
	URL   string `json:"url,omitempty" validate:"required"`
	Album string `json:"album,omitempty" validate:"required"`
}
