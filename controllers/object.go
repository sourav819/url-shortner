package controllers

type SignUpRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type AuthResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

type LoginRequest struct {
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Password    string `json:"password"`
}
