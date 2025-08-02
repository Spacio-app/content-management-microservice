package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rating struct {
	ID        string             `bson:"_id,omitempty" json:"id,omitempty"`
	ContentID primitive.ObjectID `bson:"content_id" json:"content_id"`
	Author    Author             `bson:"author" json:"author"`
	Rating    float64            `bson:"rating" json:"rating"`
	CreatedAt time.Time          `bson:"createdat,omitempty" json:"createdat,omitempty"`
	UpdatedAt time.Time          `bson:"updatedat,omitempty" json:"updatedat,omitempty"`
}
