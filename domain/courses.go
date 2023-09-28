package domain

import "time"

type CourseReq struct {
	ID                 string    `json:"ID" validate:"required"`
	Title              string    `json:"Title" validate:"required"`
	Description        string    `json:"Description" validate:"required"`
	ContentType        string    `json:"ContentType" validate:"required"`
	Author             string    `json:"Author" validate:"required"`
	CreatedAt          time.Time `json:"createdat" validate:"required"`
	UpdatedAt          time.Time `json:"updatedat" validate:"required"`
	Miniature          string    `json:"Miniature" validate:"required"`
	VideosTitles       []string  `json:"VideosTitles" validate:"required"`
	VideosDescriptions []string  `json:"VideosDescriptions" validate:"required"`
	VideosURL          []string  `json:"VideosURL" validate:"required"`
}

// Función para establecer CreatedAt y UpdatedAt antes de insertar
func (c *CourseReq) BeforeInsert() {
	currentTime := time.Now()
	c.CreatedAt = currentTime
	c.UpdatedAt = currentTime
}

// Función para actualizar UpdatedAt antes de actualizar
func (c *CourseReq) BeforeUpdate() {
	c.UpdatedAt = time.Now()
}
