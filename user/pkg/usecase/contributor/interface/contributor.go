package contributorInterfaces

import "user/pkg/domain/models"

type ContributorUsecase interface {
	Register(contributor models.Contributor) error
	Login(contributor models.LoginRequest) (string, error)

	VerifyOtp(email string, otp int) error
}
