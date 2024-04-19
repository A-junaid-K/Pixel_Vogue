package interfaces

import (
	"user/pkg/domain/models"
)

type UserRepository interface {
	Signup(user models.SignUpRequest) (models.User,error)
	// Login(user models.LoginRequest) error
	StoreOtpAndId(otp,id int)error

	CheckEmailExist(email string) (bool, error)
	CheckUserBlockOrNot(email string) (bool, error)

	GetUserById(id int)(models.User,error)
	GetUserByEmail(email string) (models.User, error)
	ValidateUser(id int)error

	DeleteUser(Id int)error
	// Create(c context.Context, user *models.User) error
	// Fetch(c context.Context) ([]models.User, error)

	// GetById(c context.Context, id string) (models.User, error)
	// Save(c context.Context, user models.User) (models.User, error)

}
