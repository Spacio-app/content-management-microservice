// class for courses implementing content interface
package models

type Courses struct {
	AbstractContent    `bson:",inline" json:",inline"`
	Miniature          string   `bson:"miniature" json:"miniature" validate:"required"`
	VideosTitles       []string `bson:"videostitle" json:"videostitle" validate:"required"`
	VideosDescriptions []string `bson:"videosdescriptions" json:"videosdescriptions" validate:"required"`
	PublicIDCloudinary []string `bson:"publicidcloudinary" json:"publicidcloudinary" validate:"required"`
	VideosURL          []string `bson:"videosurl" json:"videosurl" validate:"required"`
	// Otros campos que puedas necesitar
}
