package merchandise

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateMerchandiseParams struct {
	Name  string `form:"name" validate:"required"`
	Type  string `form:"type" validate:"required"` // TODO: use enum "Digital album" | "CD" | "T-shirt" | "Girls T-shirt" | "Patch"
	Price int    `form:"price,string" validate:"required,gte=1"`
	URL   string `form:"url" validate:"required,url"`
}

type Merchandise struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string             `json:"name,omitempty" validate:"required"`
	Type  string             `json:"type,omitempty" validate:"required"` // TODO: use enum "Digital album" | "CD" | "T-shirt" | "Girls T-shirt" | "Patch"
	Price int                `json:"price,string,omitempty" validate:"required,gte=1"`
	URL   string             `json:"url,omitempty" validate:"required,url"`
	Image string             `json:"image,omitempty" validate:"required,url"`
}
