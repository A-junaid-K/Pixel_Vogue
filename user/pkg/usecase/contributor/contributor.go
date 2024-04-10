package contributor

import (
	"errors"
	"log"
	"strconv"
	"user/pkg/domain/models"
	"user/pkg/helper"
	repository "user/pkg/repository/contributor/interfaces"
	interfaces "user/pkg/usecase/contributor/interface"
)

type contributorUseCase struct {
	contributorRepo repository.ContributorRepository
}

func NewContributorUseCase(contributorRepo repository.ContributorRepository) interfaces.ContributorUsecase {
	return &contributorUseCase{
		contributorRepo: contributorRepo,
	}
}

func (us *contributorUseCase) Register(contributor models.Contributor) error {
	log.Println("usecase : ", contributor)
	//Checking email exist or not
	exist, dberr := us.contributorRepo.ChackEmailExist(contributor.Email)
	if dberr != nil {
		return dberr
	}
	if exist {
		return errors.New("email already in use. Please login instead or use a different Email to sign up")
	}

	// Hashing user entered password
	contributor.Password = helper.Hashpassword(contributor.Password)

	registeredcontributor, err := us.contributorRepo.Register(contributor)
	if err != nil {
		return err
	}

	// Generating OTP
	otp := helper.GenerateOTP()

	// Sending generated otp to user email
	if err := helper.SendOtp(strconv.Itoa(otp), contributor.Email); err != nil {
		return err
	}

	if err := us.contributorRepo.StoreOtp(otp, registeredcontributor.Id); err != nil {
		return errors.New("failed to store otp")
	}

	return nil
}

// VerifyOtp implements interfaces.ContributorUsecase.
func (us *contributorUseCase) VerifyOtp(email string, otp int) error {

	contributor, err := us.contributorRepo.GetContributorByEmail(email)
	if err != nil {
		return err
	}
	if contributor.Email == "" {
		return errors.New("incorrect Email")
	}

	if otp == contributor.Otp {
		//	Making user validate = true
		if err := us.contributorRepo.ValidateContributor(contributor.Id); err != nil {
			return err
		}
		return nil

	} else {
		if err := us.contributorRepo.DeleteContributor(contributor.Id); err != nil {
			return err
		}
		return errors.New("incorrect OTP")
	}
}

// Login implements interfaces.UserUsecase.
func (us *contributorUseCase) Login(body models.LoginRequest) (string, error) {

	// Fetching the Contributor details
	contributor, err := us.contributorRepo.GetContributorByEmail(body.Email)
	if err != nil {
		return "", errors.New("wrong Email or Password")
	}

	//Cheking User blocked or not
	if contributor.Is_blocked {
		return "", errors.New("user blocked by Admin")
	}

	if err := helper.ComapareHashPassword(contributor.Password, body.Password); !err {
		return "", errors.New("wrong Email or password")
	}

	token, tokenerr := helper.CreateAccessTokenContributor(&contributor, "contributor")
	if tokenerr != nil {
		return "", errors.New("failed to create access token: " + tokenerr.Error())
	}

	return token, nil
}
