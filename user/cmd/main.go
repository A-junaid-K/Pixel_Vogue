// package main

// import (
// 	"fmt"
// 	"user/pkg/config"
// 	"user/pkg/di"
// )

// func main() {

// 	cfg := config.InitConfig()

// 	server, err := di.InitApi(cfg)

// 	if err != nil {
// 		fmt.Println("InitApi error: ", err)
// 	}

// 	if err := server.Start(cfg.AppPort); err != nil {
// 		fmt.Println("server error: ", err)
// 	}

// }


package main

import (
    "context"
    "log"

    pb "path/to/imageservice"

    "google.golang.org/grpc"
)

func main() {
    // Set up a connection to the gRPC server
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    // Create a new gRPC client
    client := pb.NewImageServiceClient(conn)

    // Call the UploadImage method
    contributorJWT := "your_contributor_jwt"
    imageData := []byte("your_image_data")

    req := &pb.UploadImageRequest{
        ContributorJwt: contributorJWT,
        ImageData:      imageData,
    }

    res, err := client.UploadImage(context.Background(), req)
    if err != nil {
        log.Fatalf("could not upload image: %v", err)
    }

    log.Println(res.Message)
}