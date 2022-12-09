package merchandise

type Merchandise struct {
	ID    string `json:"id" bson:"_id"`
	Name  string `json:"name"`
	Type  string `json:"type"` // TODO: use enum "Digital album" | "CD" | "T-shirt" | "Girls T-shirt" | "Patch"
	Price int    `json:"price"`
	URL   string `json:"url"`
	Image string `json:"image"`
}
