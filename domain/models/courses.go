// class for courses implementing content interface
package models

type Courses struct {
	AbstractContent
	Miniature          string   `bson:"Miniature" json:"Miniature" validate:"required"`
	VideosTitles       []string `bson:"VideosTitles" json:"VideosTitles" validate:"required"`
	VideosDescriptions []string `bson:"VideosDescriptions" json:"VideosDescriptions" validate:"required"`
	VideosURL          []string `bson:"VideosURL" json:"VideosURL" validate:"required"`
	// Otros campos que puedas necesitar
}
