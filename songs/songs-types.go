package songs

type Song struct {
	ID    string `json:"id" bson:"_id"`
	Name  string `json:"name"`
	URL   string `json:"url"`
	Album string `json:"album"`
}
