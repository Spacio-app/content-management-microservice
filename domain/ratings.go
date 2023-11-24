package domain

import (
	"time"
)

type RatingReq struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	ContentID string    `bson:"content_id" json:"content_id"`
	Author    AuthorReq `bson:"author" json:"author"`
	Rating    float64   `bson:"rating" json:"rating"`
	CreatedAt time.Time `bson:"createdat,omitempty" json:"createdat,omitempty"`
	UpdatedAt time.Time `bson:"updatedat,omitempty" json:"updatedat,omitempty"`
}

// Función para establecer CreatedAt y UpdatedAt antes de insertar
func (c *RatingReq) BeforeInsert() {
	currentTime := time.Now()
	c.CreatedAt = currentTime
	c.UpdatedAt = currentTime
}

// Función para actualizar UpdatedAt antes de actualizar
func (c *RatingReq) BeforeUpdate() {
	currentTime := time.Now()
	c.UpdatedAt = currentTime
}
