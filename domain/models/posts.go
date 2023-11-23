// class for courses implementing content interface
package models

type Posts struct {
	AbstractContent    `bson:",inline" json:",inline"`
	CreateAnnouncement bool `json:"createAnnouncement"`
	//ImagesURL          []ImageURL `bson:"imagesURL" json:"imagesURL,omitempty"`
	Blocks    []map[string]interface{} `bson:"blocks" json:"blocks,omitempty"`
	EntityMap map[string]interface{}   `bson:"entityMap" json:"entityMap,omitempty"`
	// Otros campos que puedas necesitar
	// we need dates for posts content

}
type ImageURL struct {
	ImageURL           string `bson:"imageURL" json:"imageURL,omitempty"`
	PublicIDCloudinary string `bson:"publicidcloudinary" json:"publicidcloudinary,omitempty"`
}
