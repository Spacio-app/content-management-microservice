// class for courses implementing content interface
package models

type Posts struct {
	AbstractContent    `bson:",inline" json:",inline"`
	PublicIDCloudinary []string
	ImagesURL          []string
	// Otros campos que puedas necesitar
	// we need dates for posts content

}
