package controller

import (
	"log"

	"github.com/8bury/sumup4me/internal/audio"
	"github.com/8bury/sumup4me/internal/service"
	"github.com/labstack/echo/v4"
)

type SumupController struct {
	sumupService *service.SumupService
}

func NewSumupController(api *echo.Group, sumupService *service.SumupService) *SumupController {
	log.Println("Inicializando SumupController")
	controller := &SumupController{
		sumupService: sumupService,
	}
	api.POST("/sumup/text", controller.SumupText)
	api.POST("/sumup/audio", controller.SumupAudio)
	return controller
}

func (s *SumupController) SumupText(ctx echo.Context) error {
	log.Println("Recebida requisição POST /v1/sumup/text")
	text := ctx.QueryParam("text")
	if text == "" {
		log.Println("Texto vazio recebido")
		return ctx.JSON(400, map[string]string{"error": "Texto vazio"})
	}
	log.Printf("Processando resumo para texto com %d caracteres", len(text))

	summary, err := s.sumupService.SummarizeText(text)
	if err != nil {
		log.Printf("Erro ao resumir texto: %v", err)
		return ctx.JSON(500, map[string]string{"error": "Erro ao resumir texto"})
	}

	log.Printf("Texto resumido com sucesso. Tamanho do resumo: %d caracteres", len(summary))
	return ctx.JSON(200, summary)
}

func (s *SumupController) SumupAudio(ctx echo.Context) error {
	log.Println("Recebida requisição POST /v1/sumup/audio")

	file, err := ctx.FormFile("audio")
	if err != nil {
		log.Printf("Erro ao obter arquivo de áudio: %v", err)
		return ctx.JSON(400, map[string]string{"error": "Arquivo de áudio não fornecido"})
	}

	if !audio.IsAnAudioFile(file) {
		log.Printf("Tipo de arquivo inválido: %s", file.Header.Get("Content-Type"))
		return ctx.JSON(400, map[string]string{"error": "Tipo de arquivo inválido"})
	}

	log.Printf("Processando arquivo de áudio: %s", file.Filename)
	transcription, err := s.sumupService.SumupAudio(file)
	if err != nil {
		log.Printf("Erro ao processar áudio: %v", err)
		return ctx.JSON(500, map[string]string{"error": "Erro ao processar áudio"})
	}

	log.Println("Áudio processado com sucesso")
	return ctx.JSON(200, transcription)
}
