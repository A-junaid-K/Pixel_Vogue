package middleware

import (
	"net/http"
	"user/pkg/config"
	response "user/pkg/domain/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func UserAuth(c *gin.Context) {
	cfg := config.GetConfig()
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		err := response.ErrResponse{StatusCode: http.StatusUnauthorized, Response: "Please provide your token", Error: "Empty Token"}
		c.JSON(404, err)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.AccessTokenSecret), nil
	})

	if err != nil || !token.Valid {
		resp := response.ErrResponse{StatusCode: http.StatusUnauthorized, Response: "Cannot parse autherizatoin token", Error: err.Error()}
		c.JSON(401, resp)
		return
	}

	c.Next()
}

func ContributorAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		err := response.ErrResponse{StatusCode: 401, Response: "token string is empty", Error: ""}
		c.JSON(401, err)
	}
}
