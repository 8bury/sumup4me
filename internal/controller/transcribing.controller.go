package controller

import (
	"mime/multipart"

	"github.com/8bury/sumup4me/internal/model"
	"github.com/8bury/sumup4me/internal/service"
	"github.com/labstack/echo/v4"
)

type TranscribingController struct {
	transcribingService *service.TranscribingService
}

func NewTranscribingController(api *echo.Group, transcribingService *service.TranscribingService) *TranscribingController {

	controller := &TranscribingController{
		transcribingService: transcribingService,
	}

	api.POST("/transcribe", controller.TranscribeAudio)

	return controller
}

func (c *TranscribingController) TranscribeAudio(ctx echo.Context) error {
	file, err := ctx.FormFile("audio")
	if err != nil {
		return ctx.JSON(400, model.Error{Message: "Invalid file"})
	}

	if !c.isAnAudioFile(file) {
		return ctx.JSON(400, model.Error{Message: "Invalid file type"})
	}

	transcription, err := c.transcribingService.TranscribeAudio(file)
	if err != nil {
		return ctx.JSON(500, model.Error{Message: "Failed to transcribe audio"})
	}

	return ctx.JSON(200, transcription)
}

func (c *TranscribingController) isAnAudioFile(file *multipart.FileHeader) bool {
	allowedTypes := map[string]bool{
		"audio/mpeg": true, // .mp3
		"audio/wav":  true, // .wav
		"audio/m4a":  true, // .m4a
	}
	if _, ok := allowedTypes[file.Header.Get("Content-Type")]; !ok {
		return false
	}
	return true
}
