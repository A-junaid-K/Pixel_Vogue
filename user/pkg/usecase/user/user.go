package usecase

import (
	"errors"
	"log"
	"strconv"

	"user/pkg/domain/models"
	"user/pkg/helper"
	repository "user/pkg/repository/user/interfaces"
	interfaces "user/pkg/usecase/user/interface"
)

type userUseCase struct {
	userRepo repository.UserRepository
	// contextTimeout time.Duration
}

func NewUserUseCase(userRepo repository.UserRepository) interfaces.UserUsecase {
	return &userUseCase{
		userRepo: userRepo,
		// contextTimeout: timeout,
	}
}

func (us *userUseCase) Signup(user models.SignUpRequest) error {

	//Checking email exist or not
	exist, dberr := us.userRepo.ChackEmailExist(user.Email)
	if dberr != nil {
		log.Println("db err :",dberr)
		return dberr
	}
	if exist {
		return errors.New("email already in use. Please login instead or use a different Email to sign up")
	}

	// Hashing user entered password
	user.Password = helper.Hashpassword(user.Password)
	log.Println("before:  ", user)
	signedUser, err := us.userRepo.Signup(user)
	if err != nil {
		return err
	}

	// Generating OTP
	otp := helper.GenerateOTP()

	// Sending generated otp to user email
	if err := helper.SendOtp(strconv.Itoa(otp), user.Email); err != nil {
		return err
	}

	log.Println(signedUser)
	if err := us.userRepo.StoreOtpAndId(otp, signedUser.Id); err != nil {
		return errors.New("failed to store otp")
	}

	return nil
}

// VerifyOtp implements interfaces.UserUsecase.
func (us *userUseCase) VerifyOtp(email string, otp int) error {

	user, err := us.userRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}
	if user.Email == "" {
		return errors.New("incorrect Email")
	}

	if otp == user.Otp {
		//	Making user validate = true
		if err := us.userRepo.ValidateUser(user.Id); err != nil {
			return err
		}
		return nil

	} else {
		if err := us.userRepo.DeleteUser(user.Id); err != nil {
			return err
		}
		return errors.New("incorrect OTP")
	}
}

// Login implements interfaces.UserUsecase.
func (us *userUseCase) Login(body models.LoginRequest) (string, error) {

	// Fetching the User details
	user, err := us.userRepo.GetUserByEmail(body.Email)
	if err != nil {
		return "", errors.New("wrong Email or Password")
	}
	// if user.Email == "" {
	// 	return "", errors.New("email does not exist")
	// }

	//Cheking User blocked or not
	if user.Is_blocked {
		return "", errors.New("user blocked by Admin")
	}

	if err := helper.ComapareHashPassword(user.Password, body.Password); !err {
		return "", errors.New("wrong Email or password")
	}

	token, tokenerr := helper.CreateAccessToken(&user, "user")
	if tokenerr != nil {
		return "", errors.New("failed to create access token: " + tokenerr.Error())
	}

	return token, nil
}

// func (su *userUseCase) CreateAccessToken(user *models.User, secret string, expiry int, role string) (accessToken string, err error) {
// 	return helper.CreateAccessToken(user, secret, expiry, role)
// }

// func (su *userUseCase) CreateRefreshToken(user *models.User, secret string, expiry int) (refreshToken string, err error) {
// 	return helper.CreateRefreshToken(user, secret, expiry)
// }

// func (c *userUseCase) FindAll(ctx context.Context) ([]models.User, error) {
// 	users, err := c.userRepo.Fetch(ctx)
// 	return users, err
// }

// func (u *userUseCase) FindById(c context.Context, Id string) (models.User, error) {
// 	user, err := u.userRepo.GetById(c, Id)
// 	return user, err
// }

// func (u *userUseCase) Save(c context.Context, input models.User) (models.User, error) {
// 	user, err := u.userRepo.Save(c, input)
// 	return user, err
// }

// func (su *userUseCase) Create(c context.Context, user *models.User) error {
// 	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
// 	defer cancel()
// 	return su.userRepo.Create(ctx, user)
// }

// func (su *userUseCase) GetUserByEmail(c context.Context, email string) (models.User, error) {
// 	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
// 	defer cancel()
// 	return su.userRepo.GetByEmail(ctx, email)
// }
