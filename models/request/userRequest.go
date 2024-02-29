package request

type UserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role" validate:"required,oneof=admin user"`
}

type UserUpdateRequest struct {
	Username string `json:"username"`
}

type UserUpdateEmailRequest struct {
	Email string `json:"email" validate:"required"`
}
