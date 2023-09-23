package domain

type PostReq struct {
	ID          string   `json:"ID" validate:"required"`
	Title       string   `json:"Title" validate:"required"`
	Description string   `json:"Description" validate:"required"`
	ContentType string   `json:"ContentType" validate:"required"`
	Author      string   `json:"Author" validate:"required"`
	ImagesURL   []string `json:"ImagesURL" validate:"required"`
}
