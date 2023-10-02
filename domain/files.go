package domain

import "time"

type FileReq struct {
	ID                 string    `json:"id,omitempty" bson:"_id,omitempty"`
	Title              string    `json:"Title" validate:"required"`
	Description        string    `json:"Description" validate:"required"`
	ContentType        string    `json:"ContentType" validate:"required"`
	Author             string    `json:"Author" validate:"required"`
	CreatedAt          time.Time `json:"createdat" validate:"required"`
	UpdatedAt          time.Time `json:"updatedat" validate:"required"`
	Miniature          string    `bson:"miniature" json:"miniature,omitempty"`
	PublicIDCloudinary []string  `json:"PublicIDCloudinary" validate:"required"`
	FilesURL           []string  `json:"FilesURL" validate:"required"`
}

// Función para establecer CreatedAt y UpdatedAt antes de insertar
func (c *FileReq) BeforeInsert() {
	currentTime := time.Now()
	c.CreatedAt = currentTime
	c.UpdatedAt = currentTime
	//establecer contentType
	c.ContentType = "file"
}

// Función para actualizar UpdatedAt antes de actualizar
func (c *FileReq) BeforeUpdate() {
	c.UpdatedAt = time.Now()
}
