package config

import (
	"log"
	"os"

	"github.com/8bury/sumup4me/internal/controller"
	"github.com/8bury/sumup4me/internal/dao"
	"github.com/8bury/sumup4me/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/joho/godotenv"
)

var (
	transcribingController *controller.TranscribingController
	transcribingService    *service.TranscribingService
	transcribingDao        *dao.TranscribingDao
)

func registerDependencies(api *echo.Echo, transcriptionService openai.AudioTranscriptionService) {
	v1 := api.Group("/v1")

	registerDaos(transcriptionService)
	registerServices()
	registerControllers(v1)
}

func ConfigureApi(api *echo.Echo) {
	godotenv.Load()

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is not set")
	}

	transcriptionService := openai.NewAudioTranscriptionService(option.WithAPIKey(
		apiKey,
	))

	registerDependencies(api, transcriptionService)
}

func registerDaos(transcriptionService openai.AudioTranscriptionService) {
	transcribingDao = dao.NewTranscribingDao(transcriptionService)
}
func registerServices() {
	transcribingService = service.NewTranscribingService(transcribingDao)
}
func registerControllers(api *echo.Group) {
	transcribingController = controller.NewTranscribingController(api, transcribingService)
}
