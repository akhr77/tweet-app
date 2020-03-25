package utills

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AwsS3 struct {
	Config     *Config
	Keys       AwsS3URLs
	Uploader   *s3manager.Uploader
	Downloader *s3manager.Downloader
}

type AwsS3URLs struct {
	Dir string
}

func NewAwsS3() *AwsS3 {

	config := NewConfig()

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
			Dir: "images/develop",
		},
		Uploader:   s3manager.NewUploader(sess),
		Downloader: s3manager.NewDownloader(sess),
	}
}

func (a *AwsS3) UploadImage(file string, fileName string, fileType string) (url string, err error) {

	if fileName == "" {
		return "", errors.New("fileName is required")
	}
	b64data := file[strings.IndexByte(file, ',')+1:]

	data, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		log.Fatal(err)
	}
	wb := new(bytes.Buffer)
	wb.Write(data)

	// Upload the file to S3.
	result, err := a.Uploader.Upload(&s3manager.UploadInput{
		ACL:         aws.String("public-read"),
		Body:        wb,
		Bucket:      aws.String(a.Config.Aws.S3.Bucket),
		ContentType: aws.String(fileType),
		Key:         aws.String(a.Keys.Dir + "/" + fileName),
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}

	return result.Location, nil
}

func (a *AwsS3) DownloadImage(imagePath string) (image string, err error) {

	file := &aws.WriteAtBuffer{}

	_, err = a.Downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(a.Config.Aws.S3.Bucket),
		// Key:    aws.String("images/develop/login_background.jpg"),
		Key: aws.String(imagePath),
	})

	if err != nil {
		return "", fmt.Errorf("failed to download file, %v", err)
	}
	// fmt.Printf("numBytes:%d", n)

	base64Data := base64.StdEncoding.EncodeToString(file.Bytes())

	return base64Data, nil
}
