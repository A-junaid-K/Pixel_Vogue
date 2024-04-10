package helper

import (
	"time"
	"user/pkg/config"
	"user/pkg/domain/models"

	"github.com/golang-jwt/jwt"
)

func CreateAccessToken(user *models.User, role string) (accessToken string, err error) {
	cfg := config.GetConfig()
	exp := time.Now().Add(time.Hour * time.Duration(config.GetConfig().AccessTokenExpiryHour)).Unix()
	claims := &models.JwtCustomClaims{
		Email: user.Email,
		Id:    user.Id,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token_string, err := token.SignedString([]byte(cfg.AccessTokenSecret))
	if err != nil {
		return "", err
	}
	return token_string, nil
}

func CreateRefreshToken(user *models.User) (refreshTokens string, err error) {
	cfg := config.GetConfig()
	exp := time.Now().Add(time.Hour * time.Duration(cfg.RefreshTokenExpiryHour)).Unix()
	refreshClaim := &models.JwtCustomRefreshClaim{
		Id: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim)
	refresh_token, err := token.SignedString(cfg.RefreshTokenSecret)
	if err != nil {
		return "", err
	}
	return refresh_token, err
}


//______________________________CONTRIBUTOR____________________________//


func CreateAccessTokenContributor(user *models.Contributor, role string) (accessToken string, err error) {
	cfg := config.GetConfig()
	exp := time.Now().Add(time.Hour * time.Duration(config.GetConfig().AccessTokenExpiryHour)).Unix()
	claims := &models.JwtCustomClaims{
		Email: user.Email,
		Id:    user.Id,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token_string, err := token.SignedString([]byte(cfg.AccessTokenSecret))
	if err != nil {
		return "", err
	}
	return token_string, nil
}
