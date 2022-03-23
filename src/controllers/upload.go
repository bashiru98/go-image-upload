package controllers

import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gofiber/fiber/v2"
)

type Name struct {
}

type User struct {
	Email    string
	Password string
}

func UploadFile(c *fiber.Ctx) error {

	s3Config := &aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials("KeyID", "SecretKey", ""),
	}

	s3Session := session.New(s3Config)

	uploader := s3manager.NewUploader(s3Session)
	fileheader, err := c.FormFile("image")
	if err != nil {
		return err
	}

	file, err := fileheader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	data1, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	fmt.Println("byte slice data", data1)
	fmt.Println("byte slice data file", file)

	fmt.Println("context data", file)

	input := &s3manager.UploadInput{
		Bucket: aws.String("cloudfront-bucket-2"), // bucket's name
		Key:    aws.String("myfiles/my_cat.jpg"),  // files destination location
		// Body:        bytes.NewReader(file),                   // content of the file
		ContentType: aws.String("image/jpg"), // content type
	}

	output, err := uploader.UploadWithContext(c.Context().Background(), input)
	fmt.Println("in put data", output)
	return c.JSON(data1)

}
