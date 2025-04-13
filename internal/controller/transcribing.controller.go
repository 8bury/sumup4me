package controller

import (
	"log"

	"github.com/8bury/sumup4me/internal/audio"
	"github.com/8bury/sumup4me/internal/model"
	"github.com/8bury/sumup4me/internal/service"
	"github.com/labstack/echo/v4"
)

type TranscribingController struct {
	transcribingService *service.TranscribingService
}

func NewTranscribingController(api *echo.Group, transcribingService *service.TranscribingService) *TranscribingController {
	log.Println("Inicializando TranscribingController")

	controller := &TranscribingController{
		transcribingService: transcribingService,
	}

	api.POST("/transcribe", controller.TranscribeAudio)

	return controller
}

func (c *TranscribingController) TranscribeAudio(ctx echo.Context) error {
	log.Println("Recebida requisição POST /v1/transcribe")

	file, err := ctx.FormFile("audio")
	if err != nil {
		log.Printf("Erro ao obter arquivo: %v", err)
		return ctx.JSON(400, model.Error{Message: "Invalid file"})
	}
	log.Printf("Arquivo recebido: %s (%d bytes, tipo: %s)", file.Filename, file.Size, file.Header.Get("Content-Type"))

	if !audio.IsAnAudioFile(file) {
		log.Printf("Tipo de arquivo inválido: %s", file.Header.Get("Content-Type"))
		return ctx.JSON(400, model.Error{Message: "Invalid file type"})
	}
	log.Println("Validação de tipo de arquivo realizada com sucesso")

	log.Println("Iniciando transcrição do áudio...")
	transcription, err := c.transcribingService.TranscribeAudio(file)
	if err != nil {
		log.Printf("Erro ao transcrever áudio: %v", err)
		return ctx.JSON(500, model.Error{Message: "Failed to transcribe audio " + err.Error()})
	}
	log.Println("Áudio transcrito com sucesso")

	log.Println("Enviando resposta com transcrição")
	return ctx.JSON(200, transcription)
}

