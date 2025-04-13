package service

import (
	"log"
	"mime/multipart"

	"github.com/8bury/sumup4me/internal/dao"
)

type SumupService struct {
	sumupDao            *dao.SumupDao
	transcribingService *TranscribingService
}

func NewSumupService(sumupDao *dao.SumupDao, transcribingService *TranscribingService) *SumupService {
	log.Println("Inicializando SumupService")
	return &SumupService{
		sumupDao:            sumupDao,
		transcribingService: transcribingService,
	}
}

func (s *SumupService) SummarizeText(text string) (string, error) {
	log.Printf("Solicitando resumo para texto com %d caracteres", len(text))
	summary, err := s.sumupDao.SummarizeText(text)
	if err != nil {
		log.Printf("Erro ao resumir texto: %v", err)
		return "", err
	}
	log.Printf("Resumo finalizado com sucesso. Tamanho do resumo: %d caracteres", len(summary))
	return summary, nil
}

func (s *SumupService) SumupAudio(file *multipart.FileHeader) (string, error) {
	log.Println("Iniciando transcrição e resumo do áudio")

	transcription, err := s.transcribingService.TranscribeAudio(file)
	if err != nil {
		log.Printf("Erro ao transcrever áudio: %v", err)
		return "", err
	}

	log.Printf("Transcrição concluída com sucesso. Iniciando resumo do texto")
	summary, err := s.SummarizeText(transcription.Transcription)
	if err != nil {
		log.Printf("Erro ao resumir texto: %v", err)
		return "", err
	}

	log.Println("Resumo concluído com sucesso")
	return summary, nil
}
