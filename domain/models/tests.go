// class for courses implementing content interface
package models

type Tests struct {
	AbstractContent `bson:",inline" json:",inline"`
	Questions       []Question `bson:"questions" json:"questions" validate:"required"`
	// Otros campos que puedas necesitar

}

type Question struct {
	QuestionText string   `bson:"questiontext" json:"questiontext" validate:"required"`
	Options      []Option `bson:"options" json:"options" validate:"required"`
	// Otros campos que puedas necesitar
}

type Option struct {
	OptionText string `bson:"optiontext" json:"optiontext" validate:"required"`
	IsCorrect  bool   `bson:"iscorrect" json:"iscorrect"`
	// Otros campos que puedas necesitar
}
