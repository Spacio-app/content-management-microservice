package domain

import "time"

type CourseReq struct {
	ID                 string     `json:"id,omitempty"`
	Title              string     `json:"Title" validate:"required"`
	Description        string     `json:"Description" validate:"required"`
	Author             AuthorReq  `bson:"author" json:"author" `
	Miniature          string     `json:"miniature" validate:"required"`
	PublicIDMiniature  string     `json:"publicidminiature" validate:"required"`
	ContentType        string     `json:"ContentType" validate:"required"`
	CreatedAt          time.Time  `json:"createdat,omitempty"`
	UpdatedAt          time.Time  `json:"updatedat,omitempty"`
	CreateAnnouncement bool       `json:"createAnnouncement"`
	Videos             []VideoReq `json:"videos" validate:"required"`
}
type AuthorReq struct {
	Name  string `bson:"name" json:"name"`
	Photo string `bson:"photo" json:"photo"`
	Email string `bson:"email" json:"email"`
}

type VideoReq struct {
	Title              string `json:"title" validate:"required"`
	Description        string `json:"desc" validate:"required"`
	MiniatureVideo     string `json:"miniatureVideo"`
	URL                string `json:"url" validate:"required"`
	PublicIDCloudinary string `json:"publicidcloudinary"`
	PublicIDMiniature  string `json:"publicidminiature"`
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
