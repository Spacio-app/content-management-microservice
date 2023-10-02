// class for courses implementing content interface
package models

type Files struct {
	AbstractContent    `bson:",inline" json:",inline"`
	PublicIDCloudinary []string `bson:"public_id_cloudinary" json:"public_id_cloudinary"`
	FilesURL           []string `bson:"files_url" json:"files_url"`
	// Otros campos que puedas necesitar
}
