// class for courses implementing content interface
package models

type Files struct {
	AbstractContent
	FilePath string `bson:"filePath" json:"filePath"`
	// Otros campos que puedas necesitar
}
