package videos

type Video struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name,omitempty" validate:"required"`
	URL  string `json:"url,omitempty" validate:"required"`
}
