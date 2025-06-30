package storage

import "github.com/ruiborda/go-service-common/storage/impl"

type R2Storage interface {
	UploadFile(fileData *[]byte) (fileName string, err error)
	UploadBase64File(base64File *string) (fileName string, err error)
	HeadObject(fileName string) *impl.HeadObject
	DeleteFile(fileName string) (err error)
}
