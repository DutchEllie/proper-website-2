package entity

import "time"

// Blogpost contains the name, date and such of a blogpost
// The actual post itself is hosted somewhere else in the form of an (html) document
// This is put in the path field

type BlogPost struct {
	ID       string    `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string    `json:"title" bson:"title"`
	Author   string    `json:"author" bson:"author"`
	Path     string    `json:"path" bson:"path"`
	PostDate time.Time `json:"time" bson:"time"`
}
