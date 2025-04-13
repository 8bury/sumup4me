package service

import (
	"io"
	"log"
	"mime/multipart"
	"os"

	"github.com/8bury/sumup4me/internal/dao"
	"github.com/8bury/sumup4me/internal/model"
)

type TranscribingService struct {
	TranscribingDao *dao.TranscribingDao
}

func NewTranscribingService(transcribingDao *dao.TranscribingDao) *TranscribingService {
	log.Println("Inicializando TranscribingService")
	return &TranscribingService{
		TranscribingDao: transcribingDao,
	}
}

func (transcribingService *TranscribingService) TranscribeAudio(file *multipart.FileHeader) (*model.Transcription, error) {
	log.Printf("Processando arquivo de áudio: %s", file.Filename)

	src, err := file.Open()
	if err != nil {
		log.Printf("Erro ao abrir o arquivo: %v", err)
		return nil, err
	}
	defer src.Close()
	log.Println("Arquivo de áudio aberto com sucesso")

	tempFile, err := os.CreateTemp("", "audio-*.mp3")
	if err != nil {
		log.Printf("Erro ao criar arquivo temporário: %v", err)
		return nil, err
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()
	log.Printf("Arquivo temporário criado: %s", tempFile.Name())

	_, err = io.Copy(tempFile, src)
	if err != nil {
		log.Printf("Erro ao copiar conteúdo para arquivo temporário: %v", err)
		return nil, err
	}
	log.Println("Conteúdo copiado para arquivo temporário com sucesso")

	if _, err := tempFile.Seek(0, 0); err != nil {
        log.Printf("Erro ao reposicionar ponteiro do arquivo temporário: %v", err)
        return nil, err
    }

	log.Println("Enviando arquivo para transcrição...")
	transcription, err := transcribingService.TranscribingDao.TranscribeAudio(tempFile)
	if err != nil {
		log.Printf("Erro retornado pelo serviço de transcrição: %v", err)
		return nil, err
	}

	log.Printf("Transcrição concluída com sucesso, resultado com %d caracteres", len(transcription.Transcription))
	return transcription, nil
}
