package domain

import "time"

type PostReq struct {
	Author AuthorReq `bson:"author" json:"author" validate:"required"`
	//ImagesURL          []ImageURLReq `json:"imagesURL" validate:"required"`
	CreateAnnouncement bool                     `json:"createAnnouncement"`
	Blocks             []map[string]interface{} `bson:"blocks" json:"blocks,omitempty" validate:"required"`
	EntityMap          map[string]interface{}   `bson:"entityMap" json:"entityMap,omitempty" validate:"required"`
	CreatedAt          time.Time                `json:"createdat" `
	UpdatedAt          time.Time                `json:"updatedat"`
	ContentType        string                   `json:"ContentType"`
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
