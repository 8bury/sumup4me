package service

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/8bury/sumup4me/internal/dao"
	"github.com/8bury/sumup4me/internal/model"
)

type TranscribingService struct {
	TranscribingDao *dao.TranscribingDao
}

func NewTranscribingService(transcribingDao *dao.TranscribingDao) *TranscribingService {
	return &TranscribingService{
		TranscribingDao: transcribingDao,
	}
}

func (transcribingService *TranscribingService) TranscribeAudio(file *multipart.FileHeader) (*model.Transcription, error) {

	src, err := os.Open(file.Filename)
	if err != nil {
		return nil, err
	}
	defer src.Close()

	tempFile, err := os.CreateTemp("", "audio-*")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	_, err = io.Copy(tempFile, src)
	if err != nil {
		return nil, err
	}
	return transcribingService.TranscribingDao.TranscribeAudio(tempFile)
}