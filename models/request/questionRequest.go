package request

type QuestionRequest struct {
	Question string   `json:"question" validate:"required"`
	Points   []string `json:"points" validate:"required"`
}

type QuestionUpdateRequest struct {
	Question string   `json:"question"`
	Points   []string `json:"points"`
}
