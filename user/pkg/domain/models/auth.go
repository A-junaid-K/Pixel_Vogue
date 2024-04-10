package models

import "github.com/golang-jwt/jwt"

type SignUpRequest struct {
	Name            string `json:"name,omitempty" validate:"min=3,max=20"`
	Email           string `json:"email,omitempty" validate:"email"`
	Password        string `json:"password,omitempty" validate:"min=6"`
	ConfirmPassword string `json:"confirmpassword" validate:"required,eqfield=Password"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type VerifyOtp struct {
	Email string `json:"email,omitempty" validate:"email"`
	Otp   int    `json:"otp,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email,omitempty" validate:"email"`
	Password string `json:"password,omitempty" validate:"min=6"`
}
type LoginResopnse struct {
	StatusCode int
	Token      string
}

type JwtCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaim struct {
	Id int `json:"id"`
	jwt.StandardClaims
}
