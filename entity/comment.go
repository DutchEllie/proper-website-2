package entity

import "time"

type Comment struct {
	ID       string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string    `json:"name" bson:"name"`
	Website  string    `json:"website,omitempty" bson:"website,omitempty"`
	Email    string    `json:"email,omitempty" bson:"email,omitempty"`
	Message  string    `json:"message" bson:"message"`
	PostDate time.Time `json:"time" bson:"time"`
}
