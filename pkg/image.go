package pkg

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func DownloadImageWithValidation(url string) ([]byte, string, error) {
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, "", fmt.Errorf("ошибка загрузки: %v", err)
	}
	defer resp.Body.Close()

	// Читаем первые 512 байт для определения MIME-типа
	buffer := make([]byte, 512)
	_, err = resp.Body.Read(buffer)
	if err != nil {
		return nil, "", fmt.Errorf("ошибка чтения MIME-типа: %v", err)
	}

	mimeType := http.DetectContentType(buffer)
	if !strings.HasPrefix(mimeType, "image/") {
		return nil, "", fmt.Errorf("это не изображение (MIME: %s)", mimeType)
	}

	image, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}
	return image, mimeType, nil
}

func BytesToBase64(imageData []byte) string {
	return base64.StdEncoding.EncodeToString(imageData)
}
