package domain

type CourseReq struct {
	ID                 string   `json:"ID" validate:"required"`
	Title              string   `json:"Title" validate:"required"`
	Description        string   `json:"Description" validate:"required"`
	ContentType        string   `json:"ContentType" validate:"required"`
	Author             string   `json:"Author" validate:"required"`
	Miniature          string   `json:"Miniature" validate:"required"`
	VideosTitles       []string `json:"VideosTitles" validate:"required"`
	VideosDescriptions []string `json:"VideosDescriptions" validate:"required"`
	VideosURL          []string `json:"VideosURL" validate:"required"`
}

// req interface
func (c *CourseReq) GetTitle() string {
	return c.Title
}

func (c *CourseReq) GetID() string {
	return c.ID
}
