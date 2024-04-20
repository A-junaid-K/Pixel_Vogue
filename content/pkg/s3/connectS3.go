package s3

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type ImageConfig struct {
	Client   *s3.Client
	Bucket   string
	Filename string
}

func ConnectS3(filename string) ImageConfig {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Println("failed to load sdk : ", err)
	}
	// Client = s3.NewFromConfig(cfg)
	// Bucket = "pixel-vogue"
	// Filename = "sample file"

	return ImageConfig{
		Client:   s3.NewFromConfig(cfg),
		Bucket:   "pixel-vogue",
		Filename: filename,
	}

}
