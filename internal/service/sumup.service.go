package service

import (
	"log"

	"github.com/8bury/sumup4me/internal/dao"
)

type SumupService struct {
	sumupDao *dao.SumupDao
}

func NewSumupService(sumupDao *dao.SumupDao) *SumupService {
	log.Println("Inicializando SumupService")
	return &SumupService{
		sumupDao: sumupDao,
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
