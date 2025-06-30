package impl

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"log/slog"
	"mime"
	"net/http"
)

type R2StorageImpl struct {
	bucketName      string
	accountId       string
	accessKeyId     string
	accessKeySecret string
}

func NewR2StorageImpl(
	bucketName string,
	accountId string,
	accessKeyId string,
	accessKeySecret string,
) *R2StorageImpl {
	return &R2StorageImpl{
		bucketName:      bucketName,
		accountId:       accountId,
		accessKeyId:     accessKeyId,
		accessKeySecret: accessKeySecret,
	}
}

func (this *R2StorageImpl) UploadFile(file *[]byte) (fileName string, err error) {
	fileName = ""
	detectedContentType := http.DetectContentType(*file)
	extensions, err := mime.ExtensionsByType(detectedContentType)

	if err != nil || len(extensions) == 0 {
		err = fmt.Errorf("could not detect file extension")
		slog.Error("Error detecting file extension", "error", err)
		return
	}

	fileName = fmt.Sprintf("%s%s", uuid.New().String(), extensions[0])

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(this.accessKeyId, this.accessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		slog.Error("Error loading default config", "error", err)
		return
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", this.accountId))
	})

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      &this.bucketName,
		Key:         &fileName,
		Body:        bytes.NewReader(*file),
		ContentType: &detectedContentType,
	})
	if err != nil {
		slog.Error("Error uploading file", "error", err)
		return "", err
	}
	return
}

func (this *R2StorageImpl) UploadBase64File(base64File *string) (fileName string, err error) {
	decodedData, err := base64.StdEncoding.DecodeString(*base64File)
	if err != nil {
		slog.Error("Error decoding base64 file", "error", err)
		return "", err
	}
	return this.UploadFile(&decodedData)
}

func (this *R2StorageImpl) DeleteFile(fileName string) (err error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(this.accessKeyId, this.accessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		return
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", this.accountId))
	})

	_, err = client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: &this.bucketName,
		Key:    &fileName,
	})

	return
}

type HeadObject struct {
	FileName      string
	ContentLength int64
	ContentType   string
	LastModified  int64
}

func (this *R2StorageImpl) HeadObject(fileName string) *HeadObject {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(this.accessKeyId, this.accessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		return nil
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", this.accountId))
	})

	response, err := client.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: &this.bucketName,
		Key:    &fileName,
	})
	if err != nil {
		return nil
	}

	return &HeadObject{
		FileName:      fileName,
		ContentLength: *response.ContentLength,
		ContentType:   *response.ContentType,
		LastModified:  response.LastModified.Unix(),
	}
}
