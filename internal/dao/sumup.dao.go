package dao

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
)

type SumupDao struct {
	sumupService *genai.Client
}

func NewSumupDao(sumupService *genai.Client) *SumupDao {
	log.Println("Inicializando SumupDao")
	return &SumupDao{
		sumupService: sumupService,
	}
}

func (s *SumupDao) SummarizeText(text string) (string, error) {
	log.Printf("Iniciando resumo de texto com %d caracteres", len(text))

	summarizePrompt := genai.Text("Faça que o seguinte texto seja resumido e organizado em tópicos no idioma original do texto : " + text)
	log.Println("Enviando solicitação para o modelo de IA")

	summary, err := s.sumupService.GenerativeModel("gemini-2.5-pro-exp-03-25").GenerateContent(context.Background(), summarizePrompt)
	if err != nil {
		log.Printf("Erro ao gerar conteúdo com a IA: %v", err)
		return "", err
	}

	result := fmt.Sprintf("%v", summary.Candidates[0].Content.Parts[0])
	log.Printf("Resumo obtido com sucesso. Tamanho: %d caracteres", len(result))

	return result, nil
}
