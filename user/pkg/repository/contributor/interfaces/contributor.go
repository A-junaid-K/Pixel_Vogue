package contributorInterfaces

import "user/pkg/domain/models"

type ContributorRepository interface {
	Register(contributor models.Contributor) (models.Contributor,error)
	StoreOtp(otp,id int)error

	ChackEmailExist(email string) (bool, error)
	CheckContributorBlockOrNot(email string) (bool, error)

	GetContributorByEmail(email string) (models.Contributor, error)
	ValidateContributor(id int)error

	DeleteContributor(Id int)error
}
