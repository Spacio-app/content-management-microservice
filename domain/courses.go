package domain

import "time"

type CourseReq struct {
	ID                 string       `json:"id,omitempty"`
	Title              string       `json:"Title" validate:"required"`
	Description        string       `json:"Description" validate:"required"`
	ContentType        string       `json:"ContentType" validate:"required"`
	Author             string       `json:"Author" validate:"required"`
	CreatedAt          time.Time    `json:"createdat,omitempty"`
	UpdatedAt          time.Time    `json:"updatedat,omitempty"`
	Miniature          string       `json:"Miniature" validate:"required"`
	PublicIDCloudinary []string     `json:"PublicIDCloudinary" validate:"required"`
	Videos             []VideosInfo `json:"videos" validate:"required"`
}
type VideosInfo struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	URL         string `json:"url" validate:"required"`
}

// Función para establecer CreatedAt y UpdatedAt antes de insertar
func (c *CourseReq) BeforeInsert() {
	currentTime := time.Now()
	c.CreatedAt = currentTime
	c.UpdatedAt = currentTime
	//establecer contentType
	c.ContentType = "course"
}

// Función para actualizar UpdatedAt antes de actualizar
func (c *CourseReq) BeforeUpdate() {
	currentTime := time.Now()
	c.UpdatedAt = currentTime
}
