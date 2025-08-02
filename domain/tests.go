package domain

import "time"

type TestReq struct {
	ID                 string     `json:"id" validate:"required"`
	Title              string     `json:"Title" validate:"required"`
	Description        string     `json:"Description" validate:"required"`
	ContentType        string     `json:"ContentType" validate:"required"`
	Author             AuthorReq  `json:"Author"`
	CreatedAt          time.Time  `json:"createdat" validate:"required"`
	UpdatedAt          time.Time  `json:"updatedat" validate:"required"`
	Miniature          string     `bson:"miniature" json:"miniature,omitempty"`
	CreateAnnouncement bool       `json:"createAnnouncement"`
	Questions          []Question `json:"questions" validate:"required"`
	// Otros campos que puedas necesitar
	UserAnswers []UserAnswer `json:"useranswers" validate:"required"`
	Attempts    int          `json:"attempts,omitempty"`
}

type Question struct {
	QuestionText string   `json:"questionText" validate:"required"`
	Options      []Option `json:"options" validate:"required"`
	// Otros campos que puedas necesitar
}

type Option struct {
	OptionText string `json:"optionText" validate:"required"`
	IsCorrect  bool   `json:"isCorrect" validate:"required"`
	// Otros campos que puedas necesitar
}

type UserAnswer struct {
	QuestionID string `bson:"questionid" json:"questionid" validate:"required"`
	AnswerText string `bson:"answertext" json:"answertext" validate:"required"`
	IsCorrect  bool   `bson:"iscorrect" json:"iscorrect" validate:"required"`
	// Otros campos que puedas necesitar
}
type TestResultReq struct {
	ContentID         string       `bson:"contentid" json:"contentid" validate:"required"`
	Author            AuthorReq    `bson:"userid" json:"userid" validate:"required"`
	Answers           []UserAnswer `bson:"answers" json:"answers" validate:"required"`
	Calification      float64      `bson:"calification" json:"calification" validate:"required"`
	PercentageCorrect float64      `bson:"percentagecorrect" json:"percentagecorrect" validate:"required"`
	CreatedAt         time.Time    `bson:"createdat,omitempty" json:"createdat,omitempty"`
	UpdatedAt         time.Time    `bson:"updatedat,omitempty" json:"updatedat,omitempty"`
	Attempts          int          `json:"attempts,omitempty"`
}

// Función para establecer CreatedAt y UpdatedAt antes de insertar
func (c *TestResultReq) BeforeInsert() {
	currentTime := time.Now()
	c.CreatedAt = currentTime
	c.UpdatedAt = currentTime
}

// Función para actualizar UpdatedAt antes de actualizar
func (c *TestResultReq) BeforeUpdate() {
	c.UpdatedAt = time.Now()
}
func (c *TestReq) BeforeInsert() {
	currentTime := time.Now()
	c.CreatedAt = currentTime
	c.UpdatedAt = currentTime
	//establecer contentType
	c.ContentType = "test"
}

// Función para actualizar UpdatedAt antes de actualizar
func (c *TestReq) BeforeUpdate() {
	c.UpdatedAt = time.Now()
}
