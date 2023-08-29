// models/content.go
package models

type Content struct {
	ID       string  `bson:"_id,omitempty"`
	Title    string  `bson:"titulo"`
	Content  string  `bson:"descripcion"`
	ImageURL string  `bson:"imagen"`
	VideoURL *string `bson:"video,omitempty"`
	LinkURL  string  `bson:"link"`
	// Otros campos que puedas necesitar
}
