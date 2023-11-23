// class for courses implementing content interface
package models

import "time"

type GenericContent struct {
	AbstractContent `bson:",inline" json:",inline"`
	Miniature       string                   `bson:"miniature" json:"miniature,omitempty"`
	Videos          []Video                  `json:"videos,omitempty"`
	FilesURL        []FileURL                `json:"filesURL,omitempty" validate:"required"`
	ImagesURL       []ImageURL               `json:"imagesURL,omitempty"`
	Questions       []Question               `bson:"questions" json:"questions,omitempty"`
	Feed            []Feed                   `bson:"feed" json:"feed,omitempty"`
	FeedComments    []FeedComments           `bson:"feedComments" json:"feedComments,omitempty"`
	Blocks          []map[string]interface{} `bson:"blocks" json:"blocks,omitempty"`
	EntityMap       map[string]interface{}   `bson:"entityMap" json:"entityMap"`
	// Otros campos que puedas necesitar
}

func (c *GenericContent) BeforeInsert() {
	currentTime := time.Now()
	c.CreatedAt = currentTime
	c.UpdatedAt = currentTime
}

// Función para actualizar UpdatedAt antes de actualizar
func (c *GenericContent) BeforeUpdate() {
	currentTime := time.Now()
	c.UpdatedAt = currentTime
}
