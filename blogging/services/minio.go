package services

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
)

const (
	maxFileSize = 20 * 1024 * 1024 // 10MB
)

type UploadResponse struct {
	Link string `json:"link"`
	Name string `json:"name"`
}

func InitMinio() (*minio.Client, error) {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	if endpoint == "" {
		endpoint = "localhost:9000"
	}

	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	if accessKey == "" {
		accessKey = "user"
	}

	secretKey := os.Getenv("MINIO_SECRET_KEY")
	if secretKey == "" {
		secretKey = "password"
	}
	useSSLStr := os.Getenv("MINIO_USE_SSL")
	if useSSLStr == "" {
		useSSLStr = "false"
	}
	useSSL, err := strconv.ParseBool(useSSLStr)
	if err != nil {
		log.Fatal().Err(err).Msg("Invalid value for MINIO_USE_SSL")
	}
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize MinIO client")
	}

	return minioClient, nil
}

func UploadService(client *minio.Client, bucketName, path, fileName string, files []*multipart.FileHeader, byteFile []byte) ([]UploadResponse, error) {
	var response []UploadResponse
	var fileCounter int
	newFileName := fileName
	isMultiple := false

	err := checkBucket(client, bucketName)
	if err != nil {
		return nil, err
	}

	if len(files) > 1 {
		fileCounter = 0
		isMultiple = true
	}
	if files == nil {
		fileLink := fmt.Sprintf("%s/%s", path, fileName)
		_, err = client.PutObject(context.Background(), bucketName, fileLink, bytes.NewReader(byteFile), int64(len(byteFile)), minio.PutObjectOptions{
			ContentType: "application/pdf",
		})
		response = append(response, UploadResponse{
				Link: fileLink,
				Name: newFileName,
			})

	} else {

		for _, file := range files {
			fileCounter++

			//validate file size
			if file.Size > maxFileSize {
				return nil, fmt.Errorf("file size exceeds the limit of %d bytes", maxFileSize)
			}

			reader, err := file.Open()
			if err != nil {
				return nil, fmt.Errorf("Failed to opeen file %d bytes", err)
			}
			defer reader.Close()

			if path == "user_guide" {
				newFileName = strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename))
				isMultiple = false

			}
			fileName := getFileName(path, newFileName, file, isMultiple, fileCounter)
			_, err = client.PutObject(context.Background(), bucketName, fileName, reader, file.Size, minio.PutObjectOptions{})
			if err != nil {
				return nil, fmt.Errorf("Failed to upload file %s: %v", file.Filename, err)
			}

			fileLink := fileName
			response = append(response, UploadResponse{
				Link: fileLink,
				Name: newFileName,
			})
		}
	}
	return response, nil

}

func checkBucket(client *minio.Client, bucketName string) error {
	minioLocation := os.Getenv("MINIO_LOCATION")
	exists, err := client.BucketExists(context.Background(), bucketName)
	if err != nil {
		if err := client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: minioLocation}); err != nil {
			if err == nil && exists {
				return nil
			} else {
				log.Error().Err(err).Msg("failed to check bucket existence")
				return err
			}
		}
	}
	return nil

}

func getFileName(path, newFileName string, file *multipart.FileHeader, isMultiFiles bool, fileCounter int) string {
	var fileName string

	countString := ""

	// bila multi file maka jalankan randomString
	if isMultiFiles {
		countString = fmt.Sprintf("_%d", fileCounter)
	}

	if newFileName == "" {
		currentDate := time.Now().Format("2006-01-02-15:04:05.000")
		fileName = fmt.Sprintf("%s/%s%s_%s", path, currentDate, countString, filepath.Base(file.Filename))
	} else {
		fileName = fmt.Sprintf("%s/%s%s%s", path, newFileName, countString, filepath.Ext(file.Filename))
	}

	return fileName

}
func StreamToByte(stream io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(stream)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
