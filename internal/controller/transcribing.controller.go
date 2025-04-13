package controller

import (
	"log"
	"mime/multipart"
	"path/filepath"
	"strings"

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

	if !c.isAnAudioFile(file) {
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

func (c *TranscribingController) isAnAudioFile(file *multipart.FileHeader) bool {
	// Tipos permitidos por Content-Type
	allowedTypes := map[string]bool{
		"audio/mpeg": true, // .mp3
		"audio/wav":  true, // .wav
		"audio/m4a":  true, // .m4a
	}

	// Verificar Content-Type
	contentType := file.Header.Get("Content-Type")
	if _, ok := allowedTypes[contentType]; ok {
		log.Printf("Arquivo validado pelo Content-Type: %s", contentType)
		return true
	}

	// Verificar extensão do arquivo como fallback
	allowedExtensions := map[string]bool{
		".mp3": true,
		".wav": true,
		".m4a": true,
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if _, ok := allowedExtensions[ext]; ok {
		log.Printf("Arquivo validado pela extensão: %s", ext)
		return true
	}

	return false
}
