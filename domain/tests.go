package domain

import "time"

type TestReq struct {
	Title       string     `json:"Title" validate:"required"`
	Description string     `json:"Description" validate:"required"`
	ContentType string     `json:"ContentType" validate:"required"`
	Author      string     `json:"Author" validate:"required"`
	CreatedAt   time.Time  `json:"createdat" validate:"required"`
	UpdatedAt   time.Time  `json:"updatedat" validate:"required"`
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

// Función para establecer CreatedAt y UpdatedAt antes de insertar
func (c *TestReq) BeforeInsert() {
	currentTime := time.Now()
	c.CreatedAt = currentTime
	c.UpdatedAt = currentTime
}

// Función para actualizar UpdatedAt antes de actualizar
func (c *TestReq) BeforeUpdate() {
	c.UpdatedAt = time.Now()
}
