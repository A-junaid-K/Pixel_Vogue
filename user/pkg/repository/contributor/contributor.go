package contributor

import (
	"log"

	"gorm.io/gorm"

	"user/pkg/domain/models"
	interfaces "user/pkg/repository/contributor/interfaces"
)

type ContributorRepository struct {
	DB *gorm.DB
}

func NewContributorRepository(Db *gorm.DB) interfaces.ContributorRepository {
	return &ContributorRepository{DB: Db}
}

// Signup implements interfaces.UserRepository.
func (ur *ContributorRepository) Register(contributor models.Contributor) (models.Contributor, error) {
	var registeredContributor models.Contributor
	log.Println(contributor)
	err := ur.DB.Create(&models.Contributor{FullName: contributor.FullName, Email: contributor.Email, Password: contributor.Password, Role: "contributor"}).Scan(&registeredContributor).Error

	if err != nil {
		var emptycontributor models.Contributor
		return emptycontributor, err
	}
	return registeredContributor, nil
}

func (ur *ContributorRepository) StoreOtp(otp int, id int) error {
	if err := ur.DB.Table("contributors").Where("id=?", id).Update("otp", otp).Error; err != nil {
		return err
	}
	return nil
}

func (ur *ContributorRepository) ValidateContributor(Id int) error {
	if err := ur.DB.Table("contributors").Where("id=?", Id).Update("validate", true).Error; err != nil {
		return err
	}
	return nil
}

func (ur *ContributorRepository) DeleteContributor(Id int) error {
	if err := ur.DB.Table("contributors").Where("id=?", Id).Delete(&models.Contributor{}).Error; err != nil {
		return err
	}
	return nil
}

// GetByEmail implements interfaces.ContributorRepository.
func (ur *ContributorRepository) GetContributorByEmail(email string) (models.Contributor, error) {

	var contributor models.Contributor
	if err := ur.DB.Table("contributors").Where("email=?", email).First(&contributor).Error; err != nil {
		return contributor, err
	}
	return contributor, nil
}

///////////////////////////----	CHECKS-----/////////////////////////////

func (ur *ContributorRepository) ChackEmailExist(email string) (bool, error) {
	log.Println("email: ",email)
	var count int64
	if err := ur.DB.Table("contributors").Where("email=?", email).Count(&count).Error; err != nil {
		log.Println("db err : ",err)
		return true, err
	}
	return count > 0, nil
}

func (ur *ContributorRepository) CheckContributorBlockOrNot(email string) (bool, error) {
	var is_blocked bool
	query := "SELECT is_blocked FROM contributors where email = ?"
	if err := ur.DB.Raw(query, email).Scan(&is_blocked).Error; err != nil {
		return true, err
	}
	return is_blocked, nil
}
