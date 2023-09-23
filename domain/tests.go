package domain

type TestReq struct {
	ID          string     `json:"ID" validate:"required"`
	Title       string     `json:"Title" validate:"required"`
	Description string     `json:"Description" validate:"required"`
	ContentType string     `json:"ContentType" validate:"required"`
	Author      string     `json:"Author" validate:"required"`
	Questions   []Question `json:"questions" validate:"required"`
	// Otros campos que puedas necesitar

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
