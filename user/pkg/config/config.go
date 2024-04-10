package config

import (
	"log"
	"os"
	"user"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort       string `mapstructure:"APP_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBName        string `mapstructure:"DB_NAME"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	Email         string `mapstructure:"EMAIL"`
	EmailPassword string `mapstructure:"EMAIL_PASSWORD"`

	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
}

var cfg Config

func InitConfig() Config {

	// Set the current working directory to the directory containing the .env file
	if err := os.Chdir(user.CurrentWorkingDirectory()); err != nil {
		log.Fatal("Error setting current working directory:", err)
	}

	viper.AddConfigPath("../user")
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err.Error())
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal(err.Error())
	}
	return cfg
}

func GetConfig() Config {
	return cfg
}
