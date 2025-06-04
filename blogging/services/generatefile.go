package services

import (
	"encoding/base64"
	"os"
)

func EncodeImageToBase64(path string) (string, error) {
	imgBytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(imgBytes), nil
}
