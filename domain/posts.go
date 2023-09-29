package domain

import "time"

type PostReq struct {
	Title       string    `json:"Title" validate:"required"`
	Description string    `json:"Description" validate:"required"`
	ContentType string    `json:"ContentType" validate:"required"`
	Author      string    `json:"Author" validate:"required"`
	CreatedAt   time.Time `json:"createdat" validate:"required"`
	UpdatedAt   time.Time `json:"updatedat" validate:"required"`
	ImagesURL   []string  `json:"ImagesURL" validate:"required"`
}

// Función para establecer CreatedAt y UpdatedAt antes de insertar
func (c *PostReq) BeforeInsert() {
	currentTime := time.Now()
	c.CreatedAt = currentTime
	c.UpdatedAt = currentTime
}

// Función para actualizar UpdatedAt antes de actualizar
func (c *PostReq) BeforeUpdate() {
	c.UpdatedAt = time.Now()
}
