package gigs

import "go.mongodb.org/mongo-driver/bson/primitive"

type GigFormData struct {
	Name    string `form:"name" validate:"required"`
	Venue   string `form:"venue" validate:"required"`
	Address string `form:"address" validate:"required"`
	City    string `form:"city" validate:"required"`
	Date    string `form:"date" validate:"required"`
	FBEvent string `form:"fbEvent" validate:"required,url"`
}

type Gig struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name    string             `json:"name,omitempty" validate:"required"`
	Venue   string             `json:"venue,omitempty" validate:"required"`
	Address string             `json:"address,omitempty" validate:"required"`
	City    string             `json:"city,omitempty" validate:"required"`
	Date    primitive.DateTime `json:"date,omitempty" validate:"required"`
	FBEvent string             `json:"fbEvent,omitempty" validate:"required,url"`
	Image   string             `json:"image,omitempty" validate:"required,url"`
}
