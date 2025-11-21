package dto

type RegisterInput struct {
	Name            string `json:"name" validate:"required,min=3,max=50"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
	Role            string `json:"role" validate:"required,oneof=BUYER ADMIN SELLER"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResult struct {
	SessionID string `json:"sessionID"`
}

type AuthResponse struct {
	ID    string `json:"id"`
	Role  string `json:"role"`
	Email string `json:"email"`
}
