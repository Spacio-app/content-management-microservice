// models/content.go
package models

import (
	"time"
)

// abstract content type
type AbstractContent struct {
	ID                string    `bson:"_id,omitempty" json:"id,omitempty"`
	Title             string    `bson:"title" json:"title" validate:"required"`
	Description       string    `bson:"description" json:"description" validate:"required"`
	ContentType       string    `bson:"contenttype" json:"contenttype" validate:"required"`
	Author            Author    `bson:"author" json:"author" validate:"required"`
	CreatedAt         time.Time `bson:"createdat,omitempty" json:"createdat,omitempty"`
	UpdatedAt         time.Time `bson:"updatedat,omitempty" json:"updatedat,omitempty"`
	Miniature         string    `bson:"miniature" json:"miniature"`
	PublicIDMiniature string    `bson:"publicidminiature" json:"publicidminiature"`
	AuthorID          string    `bson:"authorid" json:"authorid"`
}
type Author struct {
	Name  string `bson:"name" json:"name"`
	Photo string `bson:"photo" json:"photo"`
}
