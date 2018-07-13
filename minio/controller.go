package minio

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/ggalihpp/go-backend-ggalihpp/primary"
	minio "github.com/minio/minio-go"
)

func getBucketListController(minioClient *minio.Client) ([]minio.BucketInfo, error) {
	buckets, err := minioClient.ListBuckets()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return buckets, nil

}

// UploadFile ->> Upload a single file
func UploadFile(minioClient *minio.Client, file *multipart.FileHeader, bucketName string) (string, error) {
	xfile, _ := file.Open()

	defer xfile.Close()

	// INFORMATION ABOUT THE FILE //
	//////////////////////////////
	name := primary.GenerateBytesMask(4, true)
	contentType := file.Header["Content-Type"][0]

	fmt.Println(contentType)
	extension := filepath.Ext(file.Filename)
	name = name + extension

	_, err := minioClient.PutObject(bucketName, name, xfile, file.Size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", err

	}
	return name + " uploaded.", nil
}

func GetAllFiles(minioClient *minio.Client, bucket string) (hasil []H, err error) {

	doneCh := make(chan struct{})
	defer close(doneCh)

	isRecursive := true

	//result := v.ObjectList{}

	objectCh := minioClient.ListObjects(bucket, "", isRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			err = object.Err
			return
		}

		hasil = append(hasil, H{
			"key":           object.Key,
			"size":          object.Size,
			"last-modified": object.LastModified,
			"conten-type":   object.ContentType,
		})

		// result.Objects = append(result.Objects, object) // Dipakai sebelum nya, size lebih besar

	}

	return
}
