package domain

import "time"

type PostReq struct {
	ID          string        `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string        `json:"Title" validate:"required"`
	Description string        `json:"Description" validate:"required"`
	ContentType string        `json:"ContentType" validate:"required"`
	Author      string        `json:"Author" validate:"required"`
	CreatedAt   time.Time     `json:"createdat" validate:"required"`
	UpdatedAt   time.Time     `json:"updatedat" validate:"required"`
	Miniature   string        `bson:"miniature" json:"miniature,omitempty"`
	ImagesURL   []ImageURLReq `json:"imagesURL" validate:"required"`
}
type ImageURLReq struct {
	ImageURL           string `json:"imageURL" validate:"required"`
	PublicIDCloudinary string `json:"publicidcloudinary,omitempty" validate:"required"`
}

// Función para establecer CreatedAt y UpdatedAt antes de insertar
func (c *PostReq) BeforeInsert() {
	currentTime := time.Now()
	c.CreatedAt = currentTime
	c.UpdatedAt = currentTime
	//establecer contentType
	c.ContentType = "post"
}

// Función para actualizar UpdatedAt antes de actualizar
func (c *PostReq) BeforeUpdate() {
	c.UpdatedAt = time.Now()
}
