package main

import (
	"content/pkg/config"

	"content/pkg/di"
	"log"
)

func main() {
	cfg := config.InitConfig()

	// lis, err := net.Listen("tcp",cfg.AppPort)
	// if err != nil {
	// 	log.Fatal("could not connect with %s and got error %v", cfg.AppPort,err)
	// }

	// grpcServer := grpc.NewServer()
	// pb.RegisterTokenServiceServer(grpcServer,tokenservice{})

	server, err := di.InitApi(cfg)
	if err != nil {
		log.Println("Api initialization err : ", err)
	}

	if err := server.Start(cfg.AppPort); err != nil {
		log.Println("server error : ", err)
	}

}

// package main

// import (
// 	"context"
// 	"log"
// 	"net"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/metadata"
// 	"google.golang.org/grpc/status"
// )

// // Define the custom claims struct
// type CustomClaims struct {
// 	Username string `json:"username"`
// 	jwt.StandardClaims
// }

// // GenerateJWT generates a JWT token
// func GenerateJWT(username string) (string, error) {
// 	claims := CustomClaims{
// 		Username: username,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	signedToken, err := token.SignedString([]byte("secret-key")) // Replace "secret-key" with your actual secret key
// 	if err != nil {
// 		return "", err
// 	}

// 	return signedToken, nil
// }

// // UnaryInterceptor is a gRPC unary interceptor to validate JWT token
// func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
// 	md, ok := metadata.FromIncomingContext(ctx)
// 	if !ok {
// 		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
// 	}

// 	token := md.Get("authorization")
// 	if len(token) == 0 {
// 		return nil, status.Errorf(codes.Unauthenticated, "token is not provided")
// 	}

// 	// Extract and verify JWT token
// 	_, err := jwt.ParseWithClaims(token[0], &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("secret-key"), nil // Replace "secret-key" with your actual secret key
// 	})
// 	if err != nil {
// 		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
// 	}

// 	// If token is valid, proceed to handle the request
// 	return handler(ctx, req)
// }

// // GreeterServer is the server API for Greeter service
// type GreeterServer struct{}

// // SayHello implements the SayHello method of the Greeter service
// func (s *GreeterServer) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
// 	return &HelloResponse{Message: "Hello " + in.Name}, nil
// }

// func main() {
// 	// Start a gRPC server with the interceptor
// 	lis, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(UnaryInterceptor))
// 	RegisterGreeterServer(grpcServer, &GreeterServer{})
// 	if err := grpcServer.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }
