// class for courses implementing content interface
package models

import "time"

type Tests struct {
	AbstractContent    `bson:",inline" json:",inline"`
	CreateAnnouncement bool       `json:"createAnnouncement"`
	Questions          []Question `bson:"questions" json:"questions" validate:"required"`
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
type UserAnswer struct {
	QuestionID string `bson:"questionid" json:"questionid" validate:"required"`
	AnswerText string `bson:"answertext" json:"answertext" validate:"required"`
	IsCorrect  bool   `bson:"iscorrect" json:"iscorrect" validate:"required"`
	// Otros campos que puedas necesitar
}
type TestResult struct {
	ContentID         string       `bson:"contentid" json:"contentid" validate:"required"`
	Author            Author       `bson:"userid" json:"userid" validate:"required"`
	Answers           []UserAnswer `bson:"answers" json:"answers" validate:"required"`
	Calification      float64      `bson:"calification" json:"calification" validate:"required"`
	PercentageCorrect float64      `bson:"percentagecorrect" json:"percentagecorrect" validate:"required"`
	CreatedAt         time.Time    `bson:"createdat,omitempty" json:"createdat,omitempty"`
	UpdatedAt         time.Time    `bson:"updatedat,omitempty" json:"updatedat,omitempty"`
}
