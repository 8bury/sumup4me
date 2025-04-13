package config

import (
	"log"
	"os"

	"github.com/8bury/sumup4me/internal/controller"
	"github.com/8bury/sumup4me/internal/dao"
	"github.com/8bury/sumup4me/internal/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var (
	transcribingController *controller.TranscribingController
	transcribingService    *service.TranscribingService
	transcribingDao        *dao.TranscribingDao
)

func registerDependencies(api *echo.Echo, transcriptionService openai.AudioTranscriptionService) {
	log.Println("Registrando dependências...")
	v1 := api.Group("/v1")

	registerDaos(transcriptionService)
	registerServices()
	registerControllers(v1)
	log.Println("Dependências registradas com sucesso")
}

func ConfigureApi(api *echo.Echo) {
	log.Println("Carregando variáveis de ambiente...")
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, tentando usar variáveis do ambiente")
	} else {
		log.Println("Arquivo .env carregado com sucesso")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is not set")
	}
	log.Println("OPENAI_API_KEY encontrada")

	log.Println("Iniciando serviço de transcrição OpenAI...")
	transcriptionService := openai.NewAudioTranscriptionService(option.WithAPIKey(
		apiKey,
	),
		option.WithBaseURL(
			"https://api.openai.com/v1",
	))
	log.Println("Serviço de transcrição OpenAI inicializado")

	registerDependencies(api, transcriptionService)
}

func registerDaos(transcriptionService openai.AudioTranscriptionService) {
	log.Println("Registrando DAOs...")
	transcribingDao = dao.NewTranscribingDao(transcriptionService)
	log.Println("DAOs registrados com sucesso")
}
func registerServices() {
	log.Println("Registrando serviços...")
	transcribingService = service.NewTranscribingService(transcribingDao)
	log.Println("Serviços registrados com sucesso")
}
func registerControllers(api *echo.Group) {
	log.Println("Registrando controllers...")
	transcribingController = controller.NewTranscribingController(api, transcribingService)
	log.Println("Controllers registrados com sucesso")
}
