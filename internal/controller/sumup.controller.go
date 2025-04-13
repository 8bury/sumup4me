package controller

import (
	"log"

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
	log.Println("Rota POST /sumup/text registrada")
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
