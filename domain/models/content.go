// models/content.go
package models

import (
	"time"
)

// abstract content type
type AbstractContent struct {
	Title       string    `bson:"title" json:"title" validate:"required"`
	Description string    `bson:"description" json:"description" validate:"required"`
	ContentType string    `bson:"contenttype" json:"contenttype" validate:"required"`
	Author      string    `bson:"author" json:"author" validate:"required"`
	CreatedAt   time.Time `bson:"createdat,omitempty" json:"createdat,omitempty"`
	UpdatedAt   time.Time `bson:"updatedat,omitempty" json:"updatedat,omitempty"`
}
