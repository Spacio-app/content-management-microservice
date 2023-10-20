// class for courses implementing content interface
package models

type Files struct {
	AbstractContent `bson:",inline" json:",inline"`
	FilesURL        []FileURL `json:"filesURL" validate:"required"`
}
type FileURL struct {
	FileURL            string `json:"fileURL" validate:"required"`
	PublicIDCloudinary string `json:"publicidcloudinary" validate:"required"`
}
