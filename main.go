package main

import "C"
import (
	"context"
	"fmt"
	"time"
	"unsafe"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

//export PresignedGetObject
func PresignedGetObject(
	minioEndpointP unsafe.Pointer, length1 C.int,
	minioAccessKeyP unsafe.Pointer, length2 C.int,
	minioSecretKeyP unsafe.Pointer, length3 C.int,
	bucketNameP unsafe.Pointer, length4 C.int,
	objectNameP unsafe.Pointer, length5 C.int,
	expiry C.int) (unsafe.Pointer, C.int) {
	minioEndpoint := string(C.GoBytes(minioEndpointP, length1))
	minioAccessKey := string(C.GoBytes(minioAccessKeyP, length2))
	minioSecretKey := string(C.GoBytes(minioSecretKeyP, length3))
	bucketName := string(C.GoBytes(bucketNameP, length4))
	objectName := string(C.GoBytes(objectNameP, length5))
	// fmt.Println(minioEndpoint)
	// fmt.Println(minioAccessKey)
	// fmt.Println(minioSecretKey)
	// fmt.Println(bucketName)
	// fmt.Println(objectName)
	// fmt.Println(expiry)
	urlString := presignedGetObject_(minioEndpoint, minioAccessKey, minioSecretKey, bucketName, objectName, int(expiry))
	urlStringB := []byte(urlString)
	// fmt.Println(urlStringB)
	// fmt.Println(len(urlStringB))
	return C.CBytes(urlStringB), C.int(len(urlStringB))
}

func presignedGetObject_(
	minioEndpoint string,
	minioAccessKey string,
	minioSecretKey string,
	bucketName string,
	objectName string,
	expiry int) (url string) {
	Minio, err := minio.New(minioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioAccessKey, minioSecretKey, ""),
		Secure: false,
	})
	videoName := objectName
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	expired := time.Duration(expiry) * time.Second
	presignedURL, err := Minio.PresignedGetObject(context.Background(), bucketName, videoName, expired, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	url = presignedURL.String()
	return
}
func main() {

}
