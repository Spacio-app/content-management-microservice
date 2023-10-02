// class for courses implementing content interface
package models

import "time"

type Courses struct {
	AbstractContent    `bson:",inline" json:",inline"`
	VideosTitles       []string `bson:"videostitle" json:"videostitle" validate:"required"`
	VideosDescriptions []string `bson:"videosdescriptions" json:"videosdescriptions" validate:"required"`
	PublicIDCloudinary []string `bson:"publicidcloudinary" json:"publicidcloudinary" validate:"required"`
	VideosURL          []string `bson:"videosurl" json:"videosurl" validate:"required"`
	// Otros campos que puedas necesitar
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
