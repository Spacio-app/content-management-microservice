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
	ContentType string             `bson:"contentType" json:"contentType" validate:"required"`
	Author      string             `bson:"author" json:"author" validate:"required"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}
