// class for courses implementing content interface
package models

type GenericContent struct {
	AbstractContent `bson:",inline" json:",inline"`
	Miniature       string     `bson:"miniature" json:"miniature,omitempty"`
	Videos          []Video    `json:"videos,omitempty"`
	FilesURL        []FileURL  `json:"filesURL,omitempty" validate:"required"`
	ImagesURL       []ImageURL `json:"imagesURL,omitempty"`
	Questions       []Question `bson:"questions" json:"questions,omitempty"`
	// Otros campos que puedas necesitar
}
