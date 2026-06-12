package model

type RegisterUserRequest struct {
	Email           string  `json:"email" validate:"required,email"`
	Password        string  `json:"password" validate:"required,min=6,max=100"`
	ConfirmPassword string  `json:"confirm_password" validate:"required,eqfield=Password"`
	FirstName       string  `json:"first_name" validate:"required,alpha"`
	LastName        *string `json:"last_name,omitempty" validate:"omitempty,alpha"`
}

type RegisterUserResponse struct {
	Email     string  `json:"email"`
	FirstName string  `json:"first_name"`
	LastName  *string `json:"last_name,omitempty"`
	CreatedAt string  `json:"created_at"`
}
