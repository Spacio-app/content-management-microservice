// class for courses implementing content interface
package models

import (
	"time"
)

type Courses struct {
	AbstractContent    `bson:",inline" json:",inline"`
	PublicIDCloudinary []string     `bson:"publicidcloudinary" json:"publicidcloudinary" validate:"required"`
	Videos             []VideosInfo `bson:"videos" json:"videos" validate:"required"`
	// Otros campos que puedas necesitar
}
type VideosInfo struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	URL         string `json:"url" validate:"required"`
}

func (c *Courses) BeforeInsert() {
	currentTime := time.Now()
	c.CreatedAt = currentTime
	c.UpdatedAt = currentTime
	//establecer contentType
	c.ContentType = "course"
}

// Función para actualizar UpdatedAt antes de actualizar
func (c *Courses) BeforeUpdate() {
	currentTime := time.Now()
	c.UpdatedAt = currentTime
}
