package usecase

import (
	"content/pkg/config"
	"content/pkg/domain/models"
	repointerface "content/pkg/repository/interfaces"
	interfaces "content/pkg/usecase/interfaces"
	"errors"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type ImageUsecase struct {
	ImageRepo repointerface.ImageRepository
}

func NewImageUsecase(imagerepo repointerface.ImageRepository) interfaces.ImageUsecase {
	return &ImageUsecase{
		ImageRepo: imagerepo,
	}
}

func (us *ImageUsecase) UploadImage(image multipart.File, head multipart.FileHeader, imagedetails models.ImageDetails) error {

	cfg := config.GetConfig()
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(cfg.AwsRegion),
		Credentials: credentials.NewStaticCredentials(cfg.AwsAccessKey, cfg.AwsSecretAccessKey, ""),
	}))

	filename_slice := strings.Split(head.Filename, ".")
	ext := filename_slice[len(filename_slice)-1]
	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(cfg.AwsBucket),
		Key:    aws.String("images/" + strconv.Itoa(imagedetails.ContributorId) + "." + ext),
		// Key:  &head.Filename,
		ACL:  aws.String("public-read"),
		Body: image,
	})

	if err != nil {
		return errors.New("failed to upload image in s3 bucket : " + err.Error())
	}

	if err := us.ImageRepo.UploadImage(result.Location, imagedetails); err != nil {
		return err
	}

	return nil
}
