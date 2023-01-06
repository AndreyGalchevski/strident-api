package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Doc      `bson:",inline"`
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

type Gig struct {
	Doc     `bson:",inline"`
	Name    string             `json:"name,omitempty" bson:"name" form:"name" validate:"required"`
	Venue   string             `json:"venue,omitempty" bson:"venue" form:"venue" validate:"required"`
	Address string             `json:"address,omitempty" bson:"address" form:"address" validate:"required"`
	City    string             `json:"city,omitempty" bson:"city" form:"city" validate:"required"`
	Date    primitive.DateTime `json:"date,omitempty" bson:"date" form:"date" validate:"required"`
	FBEvent string             `json:"fbEvent,omitempty" bson:"fbEvent" form:"fbEvent" validate:"required,url"`
	Image   string             `json:"image,omitempty" bson:"image,omitempty"`
}

type Lyric struct {
	Doc  `bson:",inline"`
	Name string `json:"name,omitempty" bson:"name" form:"name" validate:"required"`
	Text string `json:"text,omitempty" bson:"text" form:"text" validate:"required"`
}

type Member struct {
	Doc        `bson:",inline"`
	Name       string `json:"name,omitempty" bson:"name" form:"name" validate:"required"`
	Instrument string `json:"instrument,omitempty" bson:"instrument" form:"instrument" validate:"required"`
	Image      string `json:"image,omitempty" bson:"image,omitempty"`
}

type Merchandise struct {
	Doc   `bson:",inline"`
	Name  string `json:"name,omitempty" bson:"name" form:"name" validate:"required"`
	Type  string `json:"type,omitempty" bson:"type" form:"type" validate:"required"` // TODO: use enum "Digital album" | "CD" | "T-shirt" | "Girls T-shirt" | "Patch"
	Price int    `json:"price,omitempty" bson:"price" form:"price" validate:"required,gte=1"`
	URL   string `json:"url,omitempty" bson:"url" form:"url" validate:"required,url"`
	Image string `json:"image,omitempty" bson:"image,omitempty"`
}

type Song struct {
	Doc   `bson:",inline"`
	Name  string `json:"name,omitempty" bson:"name" form:"name" validate:"required"`
	URL   string `json:"url,omitempty" bson:"url" form:"url" validate:"required,url"`
	Album string `json:"album,omitempty" bson:"album" form:"album" validate:"required"`
}

type Video struct {
	Doc  `bson:",inline"`
	Name string `json:"name,omitempty" bson:"name" form:"name" validate:"required"`
	URL  string `json:"url,omitempty" bson:"url" form:"url" validate:"required,url"`
}

type Database struct {
	Gigs        *Collection[*Gig]
	Lyrics      *Collection[*Lyric]
	Members     *Collection[*Member]
	Merchandise *Collection[*Merchandise]
	Songs       *Collection[*Song]
	Videos      *Collection[*Video]
	Users       *Collection[*User]
}

func GetCollectionNew[T Document](collectionName string) *Collection[T] {
	return &Collection[T]{GetDBClient().Database("main").Collection(collectionName)}
}

func GetDB() Database {
	return Database{
		Gigs:        GetCollectionNew[*Gig]("gigs"),
		Lyrics:      GetCollectionNew[*Lyric]("lyrics"),
		Members:     GetCollectionNew[*Member]("members"),
		Merchandise: GetCollectionNew[*Merchandise]("merchandise"),
		Songs:       GetCollectionNew[*Song]("songs"),
		Videos:      GetCollectionNew[*Video]("videos"),
		Users:       GetCollectionNew[*User]("users"),
	}
}
