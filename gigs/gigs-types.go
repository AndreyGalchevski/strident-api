package gigs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Gig struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name    string             `json:"name,omitempty" validate:"required"`
	Venue   string             `json:"venue,omitempty" validate:"required"`
	Address string             `json:"address,omitempty" validate:"required"`
	City    string             `json:"city,omitempty" validate:"required"`
	Date    string             `json:"date,omitempty" validate:"required"`
	FBEvent string             `json:"fbEvent,omitempty" validate:"required"`
	Image   string             `json:"image,omitempty" validate:"required"`
}
