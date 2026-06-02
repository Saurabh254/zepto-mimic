package models

type LoginRequest struct {
	Email    string `json:"email" example:"test@example.com"`
	Password string `json:"password" example:"secret123"`
}

type LoginResponse struct {
	Token string `json:"token" example:"jwt-token"`
}
