package models

import "time"

type Feed struct {
	AbstractContent `bson:",inline" json:",inline"`
	ContentID       string         `bson:"content_id" json:"content_id"`
	Comments        []FeedComments `bson:"comments" json:"comments"`
}
type FeedComments struct {
	Author    Author    `bson:"author" json:"author"`
	Comment   string    `bson:"comment" json:"comment"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func (c *Feed) BeforeInsert() {
	currentTime := time.Now()
	c.CreatedAt = currentTime
	c.UpdatedAt = currentTime
}

// Función para actualizar UpdatedAt antes de actualizar
func (c *Feed) BeforeUpdate() {
	currentTime := time.Now()
	c.UpdatedAt = currentTime
}
func (c *FeedComments) BeforeInsert() {
	currentTime := time.Now()
	c.CreatedAt = currentTime
	c.UpdatedAt = currentTime
}

// Función para actualizar UpdatedAt antes de actualizar
func (c *FeedComments) BeforeUpdate() {
	currentTime := time.Now()
	c.UpdatedAt = currentTime
}
