package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeedReq struct {
	ID          string            `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string            `bson:"title" json:"title" validate:"required"`
	Description string            `bson:"description" json:"description" validate:"required"`
	ContentType string            `bson:"contenttype" json:"contenttype" validate:"required"`
	Author      AuthorReq         `bson:"author" json:"author" validate:"required"`
	CreatedAt   time.Time         `bson:"createdat,omitempty" json:"createdat,omitempty"`
	UpdatedAt   time.Time         `bson:"updatedat,omitempty" json:"updatedat,omitempty"`
	AuthorID    string            `bson:"authorid" json:"authorid"`
	ContentID   string            `bson:"content_id" json:"content_id"`
	AuthorPhoto string            `bson:"author_photo" json:"author_photo"`
	Miniature   string            `bson:"miniature" json:"miniature"`
	Comments    []FeedCommentsReq `bson:"comments" json:"comments"`
}
type FeedCommentsReq struct {
	CommentID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Author    AuthorReq          `bson:"author" json:"author"`
	Comment   string             `bson:"comment" json:"comment"`
	ContentID string             `bson:"content_id" json:"content_id"`
	CreatedAt time.Time          `bson:"createdat,omitempty" json:"createdat,omitempty"`
	UpdatedAt time.Time          `bson:"updatedat,omitempty" json:"updatedat,omitempty"`
}

// Función para establecer CreatedAt y UpdatedAt antes de insertar
func (c *FeedReq) BeforeInsert() {
	currentTime := time.Now()
	c.CreatedAt = currentTime
	c.UpdatedAt = currentTime
	//establecer contentType
	c.ContentType = "feed"
}

// Función para actualizar UpdatedAt antes de actualizar
func (c *FeedReq) BeforeUpdate() {
	currentTime := time.Now()
	c.UpdatedAt = currentTime
}
func (c *FeedCommentsReq) BeforeInsert() {
	currentTime := time.Now()
	c.CreatedAt = currentTime
	c.UpdatedAt = currentTime
}
func (c *FeedCommentsReq) BeforeUpdate() {
	currentTime := time.Now()
	c.UpdatedAt = currentTime
}
