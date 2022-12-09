package merchandise

type Merchandise struct {
	ID    string `json:"id" bson:"_id"`
	Name  string `json:"name,omitempty" validate:"required"`
	Type  string `json:"type,omitempty" validate:"required"` // TODO: use enum "Digital album" | "CD" | "T-shirt" | "Girls T-shirt" | "Patch"
	Price int    `json:"price,omitempty" validate:"required"`
	URL   string `json:"url,omitempty" validate:"required"`
	Image string `json:"image,omitempty" validate:"required"`
}
