// models/content.go
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// abstract content type
type AbstractContent struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title" validate:"required"`
	Description string             `bson:"description" json:"description" validate:"required"`
	ContentType string             `bson:"contenttype" json:"contenttype" validate:"required"`
	Author      string             `bson:"author" json:"author" validate:"required"`
	CreatedAt   time.Time          `bson:"createdat" json:"createdat" validate:"required"`
	UpdatedAt   time.Time          `bson:"updatedat" json:"updatedat" validate:"required"`
}
