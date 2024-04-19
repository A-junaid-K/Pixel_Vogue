package interfaces

import (
	"user/pkg/domain/models"
)

type UserUsecase interface {
	Signup(user models.SignUpRequest) error
	VerifyOtp(email string, otp int) error
	Login(user models.LoginRequest) (string, error)

	GetUserById(id int)(models.User,error)

	// Create(c context.Context, user *models.User) error
	// GetUserByEmail(c context.Context, email string) (models.User, error)
	// CreateAccessToken(user *models.User, secret string, expiry int,role string) (accessToken string, err error)
	// CreateRefreshToken(user *models.User, secret string, expiry int) (refreshToken string, err error)
}
