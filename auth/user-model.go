package auth

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty"`
	Username string             `json:"name,omitempty" validate:"required"`
	Password string             `json:"location,omitempty" validate:"required"`
}
