package dto

type UserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
