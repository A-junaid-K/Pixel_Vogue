package middleware

import (
	"content/pkg/config"
	"content/pkg/domain/response"
	pb "content/pkg/grpc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ContributorAuth(c *gin.Context, t *pb.TokenRequest) {
	cfg := config.GetConfig()
	tokenstring := t.ContributorToken

	if tokenstring == "" {
		err := response.ErrResponse{StatusCode: http.StatusUnauthorized, Response: "Please provide your token", Error: "Empty Token"}
		c.JSON(404, err)
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenstring, func(t *jwt.Token) (interface{}, error) {
		return cfg.ContributorSecreteAccessToken, nil
	})

	if err != nil {
		resp := response.ErrResponse{StatusCode: http.StatusUnauthorized, Response: "Cannot parse autherizatoin token", Error: err.Error()}
		c.JSON(401, resp)
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		resp := response.ErrResponse{StatusCode: 401, Response: "Invalid Authorization token"}
		c.JSON(http.StatusUnauthorized, resp)
		c.Abort()
		return
	}

	if role := claims["role"]; role != "contributor" {
		resp := response.ErrResponse{StatusCode: 403, Response: "UnAuthorized Access"}
		c.JSON(http.StatusForbidden, resp)
		c.Abort()
		return
	}

	id := claims["id"].(int)
	if id == 0 {
		resp := response.ErrResponse{StatusCode: 403, Response: "Something wrong in token"}
		c.JSON(http.StatusForbidden, resp)
		c.Abort()
		return
	}

	c.Set("id",id)
	c.Next()
}
