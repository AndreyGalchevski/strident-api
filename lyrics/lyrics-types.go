package lyrics

type Lyric struct {
	ID    string `json:"id" bson:"_id"`
	Name  string `json:"name"`
	Text  string `json:"text"`
	Album string `json:"album"`
}
