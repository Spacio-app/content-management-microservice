// class for courses implementing content interface
package models

type Files struct {
	AbstractContent `bson:",inline" json:",inline"`
	FilePath        string `bson:"filepath" json:"filepath"`
	// Otros campos que puedas necesitar
}
