package models

type BFFSinginUserRequest struct {
	Email    string `json:"email" example:"arijit@gmail.com" validate:"required,Email"`
	Password string `json:"password" example:"Ari123Jit@" validate:"required,min=8,max=20"`
}

type BFFSigninUserResponse struct {
	Username    string `json:"username" example:"Arijit"`
	Email       string `json:"email" example:"arijit@gmail.com"`
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"`
}

type SuccessAPIResponse struct {
	Message string                `json:"message" example:"User logged in successfully"`
	Data    BFFSigninUserResponse `json:"data"`
}
