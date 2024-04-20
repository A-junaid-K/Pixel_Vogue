package repository

import (
	"log"
	"user/pkg/domain/models"
	"user/pkg/repository/user/interfaces"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(Db *gorm.DB) interfaces.UserRepository {
	return &UserRepository{DB: Db}
}

// Signup implements interfaces.UserRepository.
func (ur *UserRepository) Signup(user models.SignUpRequest) (models.User, error) {
	log.Println("sssss: ", user)
	var signedUpUser models.User

	// query := "INSERT INTO users (name, email, password,created_at)Values($1,$2,$3,$4)"
	// err := ur.DB.Exec(query, user.Name, user.Email, user.Password, time.Now()).Error

	err := ur.DB.Create(&models.User{Email: user.Email, Password: user.Password}).Scan(&signedUpUser).Error

	if err != nil {
		var emptyuser models.User
		log.Println("sign up error: ", err)
		return emptyuser, err
	}

	return signedUpUser, nil
}

// StoreOtpAndId implements interfaces.UserRepository.
func (ur *UserRepository) StoreOtpAndId(otp int, id int) error {
	if err := ur.DB.Table("users").Where("id=?", id).Update("otp", otp).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) ValidateUser(Id int) error {
	if err := ur.DB.Table("users").Where("id=?", Id).Update("validate", true).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) DeleteUser(Id int) error {
	if err := ur.DB.Table("users").Where("id=?", Id).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}

// GetByEmail implements interfaces.UserRepository.
func (ur *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := ur.DB.Table("users").Where("email=?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserById(id int)(models.User,error){
	var user models.User
	if err := ur.DB.Table("users").Where("id=?",id).Scan(&user).Error; err != nil{
		return user,err
	}
	return user, nil
}

///////////////////////////----	CHECKS-----/////////////////////////////

func (ur *UserRepository) CheckEmailExist(email string) (bool, error) {
	log.Println("before :",email)
	var count int64
	if err := ur.DB.Table("users").Where("email=?", email).Count(&count).Error; err != nil {
		log.Println("db email err: ",err)
		return true, err
	}
	return count > 0, nil
}

// CheckUserBlockOrNot implements interfaces.UserRepository.
func (ur *UserRepository) CheckUserBlockOrNot(email string) (bool, error) {
	var is_blocked bool
	query := "SELECT is_blocked FROM users where email = ?"
	if err := ur.DB.Raw(query, email).Scan(&is_blocked).Error; err != nil {
		return true, err
	}
	return is_blocked, nil
}
