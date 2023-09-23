// class for courses implementing content interface
package models

type Tests struct {
	AbstractContent
	Questions []Question `bson:"questions" json:"questions" validate:"required"`
	// Otros campos que puedas necesitar

}

type Question struct {
	QuestionText string   `bson:"questionText" json:"questionText" validate:"required"`
	Options      []Option `bson:"options" json:"options" validate:"required"`
	// Otros campos que puedas necesitar
}

type Option struct {
	OptionText string `bson:"optionText" json:"optionText" validate:"required"`
	IsCorrect  bool   `bson:"isCorrect" json:"isCorrect"`
	// Otros campos que puedas necesitar
}
