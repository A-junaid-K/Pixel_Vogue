package db

import (
	"fmt"
	"content/pkg/config"
	"content/pkg/domain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Err error
)

func ConnectdB(c config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort)
	DB, Err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if Err != nil {
		return nil, Err
	}

	DB.AutoMigrate(
		&models.Image{},
	)
	return DB, nil
}
