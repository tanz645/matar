package mediaService

import (
	"bytes"
	"fmt"
	"matar/configs"
	"mime/multipart"
	"text/template"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var key = configs.GetEnvVar("SPACE_BUCKET_KEY")
var secret = configs.GetEnvVar("SPACE_BUCKET_SECRET")

func UploadFiles(
	file *multipart.FileHeader,
	bucket string,
	acl string,
	prefix string,
	metaData map[string]*string) (string, error) {

	if acl == "" {
		acl = ACCESS_LEVEL_PUBLIC
	}
	if bucket == "" {
		bucket = configs.GetEnvVar("SPACE_BUCKET_NAME")
	}
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:    aws.String(configs.GetEnvVar("SPACE_BUCKET_URL")),
		Region:      aws.String(configs.GetEnvVar("SPACE_BUCKET_REGION")),
	}

	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)
	fileContents, err := file.Open()
	if err != nil {
		return "", err
	}
	object := s3.PutObjectInput{
		Bucket:   aws.String(bucket),
		Key:      aws.String(prefix + file.Filename),
		Body:     fileContents,
		ACL:      aws.String(acl),
		Metadata: metaData,
	}
	_, putError := s3Client.PutObject(&object)
	fileContents.Close()
	if putError != nil {
		fmt.Println(putError.Error())
	}
	uploadedURL := UploadedURL{bucket, configs.GetEnvVar("SPACE_BUCKET_REGION"), "digitaloceanspaces.com", prefix, file.Filename}
	url, err := template.New("uploadedURL").Parse("https://{{.Bucket}}.{{.Region}}.cdn.{{.Host}}/{{.Prefix}}{{.FileName}}")
	if err != nil {
		return "", err
	}
	var doc bytes.Buffer
	err = url.Execute(&doc, uploadedURL)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return doc.String(), nil
}

func DeleteFile(
	filePath string,
	bucket string,
	acl string) error {
	if acl == "" {
		acl = ACCESS_LEVEL_PUBLIC
	}
	if bucket == "" {
		bucket = configs.GetEnvVar("SPACE_BUCKET_NAME")
	}
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:    aws.String(configs.GetEnvVar("SPACE_BUCKET_URL")),
		Region:      aws.String(configs.GetEnvVar("SPACE_BUCKET_REGION")),
	}
	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filePath),
	}
	_, err := s3Client.DeleteObject(input)
	if err != nil {
		return err
	}
	return nil
}
