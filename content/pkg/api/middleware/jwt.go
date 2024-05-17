package middleware

import (
	"content/pkg/config"
	"content/pkg/domain/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ContributorAuth(c *gin.Context) {
	cfg := config.GetConfig()
	tokenstring := c.GetHeader("Authorization")

	if tokenstring == "" {
		err := response.ErrResponse{StatusCode: http.StatusUnauthorized, Response: "Please provide your token", Error: "Empty Token"}
		c.JSON(404, err)
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenstring, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.ContributorAccessToken), nil
	})

	if err != nil {
		resp := response.ErrResponse{StatusCode: http.StatusUnauthorized, Response: "Cannot parse autherization token", Error: err.Error()}
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

	id := claims["id"].(float64)
	if id == 0 {
		resp := response.ErrResponse{StatusCode: 403, Response: "Something wrong in token"}
		c.JSON(http.StatusForbidden, resp)
		c.Abort()
		return
	}
	uid := int(id)
	c.Set("id", uid)
	// c.Next()
}

// package interceptor

// import (
//     "context"
//     "github.com/grpc-ecosystem/go-grpc-middleware"
//     "github.com/grpc-ecosystem/go-grpc-middleware/auth"
//     "google.golang.org/grpc"
// )

// func AuthInterceptor(ctx context.Context) (context.Context, error) {
//     // Extract JWT token from context metadata
//     token, err := grpc_auth.AuthFromMD(ctx, "bearer")
//     if err != nil {
//         return nil, err
//     }

//     // Verify token and extract contributor information
//     contributor, err := VerifyTokenAndGetContributor(token)
//     if err != nil {
//         return nil, err
//     }

//     // Attach contributor information to context
//     ctx = context.WithValue(ctx, "contributor", contributor)

//     return ctx, nil
// }

// func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
//     return grpc_auth.UnaryServerInterceptor(AuthInterceptor)
// }
