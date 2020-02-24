package infrastructure

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AwsS3 struct {
	Config   *Config
	Keys     AwsS3URLs
	Uploader *s3manager.Uploader
}

type AwsS3URLs struct {
	Prefix string
}

func NewAwsS3() *AwsS3 {

	config := NewConfig()

	// The session the S3 Uploader will use
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
				AccessKeyID:     config.Aws.S3.AccessKeyID,
				SecretAccessKey: config.Aws.S3.SecretAccessKey,
			}),
			Region: aws.String(config.Aws.S3.Region),
		},
	}))

	return &AwsS3{
		Config: config,
		Keys: AwsS3URLs{
			Prefix: "favpic/images",
		},
		// Create an uploader with the session and default options
		Uploader: s3manager.NewUploader(sess),
	}
}

func (a *AwsS3) UploadImage(file string, fileName string, extension string) (url string, err error) {

	if fileName == "" {
		return "", errors.New("fileName is required")
	}

	var contentType string

	switch extension {
	case "jpg":
		contentType = "image/jpeg"
	case "jpeg":
		contentType = "image/jpeg"
	case "gif":
		contentType = "image/gif"
	case "png":
		contentType = "image/png"
	default:
		return "", errors.New("this extension is invalid")
	}

	data, _ := base64.StdEncoding.DecodeString(file)
	wb := new(bytes.Buffer)
	wb.Write(data)

	// Upload the file to S3.
	result, err := a.Uploader.Upload(&s3manager.UploadInput{
		ACL:         aws.String("public-read"),
		Body:        wb,
		Bucket:      aws.String(a.Config.Aws.S3.Bucket),
		ContentType: aws.String(contentType),
		Key:         aws.String(a.Keys.Prefix + "/" + fileName + "." + extension),
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}

	return result.Location, nil
}
