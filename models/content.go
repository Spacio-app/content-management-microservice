// models/content.go
package models

import "time"

type AbstractContent struct {
	ID          string    `bson:"_id,omitempty" validate:"required"`
	Title       string    `bson:"titulo" validate:"required"`
	Description string    `bson:"descripcion" validate:"required"`
	ContentType string    `bson:"tipo_contenido" validate:"required"`
	Author      string    `bson:"autor" validate:"required"`
	CreatedAt   time.Time `bson:"fecha_creacion" validate:"required"`
	UpdatedAt   time.Time `bson:"fecha_actualizacion" validate:"required"`
}
