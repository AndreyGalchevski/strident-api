package members

type Member struct {
	ID         string `json:"id" bson:"_id"`
	Name       string `json:"name"`
	Instrument string `json:"instrument"`
	Image      string `json:"image"`
}
