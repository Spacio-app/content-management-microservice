// class for courses implementing content interface
package models

type GenericContent struct {
	AbstractContent `bson:",inline" json:",inline"`
	Miniature       string         `bson:"miniature" json:"miniature,omitempty"`
	Videos          []Video        `json:"videos,omitempty"`
	FilesURL        []FileURL      `json:"filesURL,omitempty" validate:"required"`
	ImagesURL       []ImageURL     `json:"imagesURL,omitempty"`
	Questions       []Question     `bson:"questions" json:"questions,omitempty"`
	Feed            []Feed         `bson:"feed" json:"feed,omitempty"`
	FeedComments    []FeedComments `bson:"feedComments" json:"feedComments,omitempty"`
	// Otros campos que puedas necesitar
}
