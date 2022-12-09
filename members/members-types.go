package members

type Member struct {
	ID         string `json:"id" bson:"_id"`
	Name       string `json:"name,omitempty" validate:"required"`
	Instrument string `json:"instrument,omitempty" validate:"required"`
	Image      string `json:"image,omitempty" validate:"required"`
}
