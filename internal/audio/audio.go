package audio

import (
	"log"
	"mime/multipart"
	"path/filepath"
	"strings"
)

func IsAnAudioFile(file *multipart.FileHeader) bool {
	const (
		contentTypeHeader = "Content-Type"
	)
	allowedTypes := map[string]struct{}{
		"audio/mpeg": {},
		"audio/wav":  {},
		"audio/m4a":  {},
		"audio/mp4": {},
		"audio/aac":  {},
	}

	allowedExtensions := map[string]struct{}{
		".mp3": {},
		".wav": {},
		".m4a": {},
		".mp4": {},
		".aac": {},
	}

	contentType := file.Header.Get(contentTypeHeader)
	if _, ok := allowedTypes[contentType]; ok {
		log.Printf("Arquivo validado pelo Content-Type: %s", contentType)
		return true
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if _, ok := allowedExtensions[ext]; ok {
		log.Printf("Arquivo validado pela extensão: %s", ext)
		return true
	}

	log.Printf("Arquivo inválido: Content-Type '%s', Extensão '%s'", contentType, ext)
	return false
}
