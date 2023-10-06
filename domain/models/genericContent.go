// class for courses implementing content interface
package models

type GenericContent struct {
	AbstractContent    `bson:",inline" json:",inline"`
	Miniature          string     `bson:"miniature" json:"miniature,omitempty"`
	VideosTitles       []string   `bson:"videostitle" json:"videostitle,omitempty"`
	VideosDescriptions []string   `bson:"videosdescriptions" json:"videosdescriptions,omitempty" `
	PublicIDCloudinary []string   `bson:"publicidcloudinary" json:"publicidcloudinary,omitempty"`
	VideosURL          []string   `bson:"videosurl" json:"videosurl,omitempty"`
	FilePath           string     `bson:"filepath" json:"filepath,omitempty"`
	ImagesURL          []string   `bson:"imagesurl" json:"imagesurl,omitempty"`
	Questions          []Question `bson:"questions" json:"questions,omitempty"`
	// Otros campos que puedas necesitar
}
