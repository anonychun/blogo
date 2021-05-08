package model

type AuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
