package request

type AnswerRequest struct {
	UserId  int                      `json:"user_id" validate:"required,number"`
	Answers []map[string]interface{} `json:"answers" validate:"required"`
}

type AnswerUpdateRequest struct {
	UserId  int                      `json:"user_id" validate:"number"`
	Answers []map[string]interface{} `json:"answers"`
}
