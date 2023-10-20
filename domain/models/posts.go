// class for courses implementing content interface
package models

type Posts struct {
	AbstractContent `bson:",inline" json:",inline"`
	ImagesURL       []ImageURL `bson:"imagesURL" json:"imagesURL,omitempty"`
	// Otros campos que puedas necesitar
	// we need dates for posts content

}
type ImageURL struct {
	ImageURL           string `bson:"imageURL" json:"imageURL,omitempty"`
	PublicIDCloudinary string `bson:"publicidcloudinary" json:"publicidcloudinary,omitempty"`
}
