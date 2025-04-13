package config

import (
	"context"
	"log"
	"os"

	"github.com/8bury/sumup4me/internal/controller"
	"github.com/8bury/sumup4me/internal/dao"
	"github.com/8bury/sumup4me/internal/service"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/openai/openai-go"
	openAIOption "github.com/openai/openai-go/option"
	geminiOption "google.golang.org/api/option"
)

var (
	transcribingController *controller.TranscribingController
	transcribingService    *service.TranscribingService
	transcribingDao        *dao.TranscribingDao

	sumupController *controller.SumupController
	sumupService    *service.SumupService
	sumupDao        *dao.SumupDao
)

func registerDependencies(api *echo.Echo, transcriptionService openai.AudioTranscriptionService, sumupService *genai.Client) {
	log.Println("Registrando dependências...")
	v1 := api.Group("/v1")

	registerDaos(transcriptionService, sumupService)
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

	apiKeyOpenAI := os.Getenv("OPENAI_API_KEY")
	if apiKeyOpenAI == "" {
		log.Fatal("OPENAI_API_KEY environment variable is not set")
	}
	log.Println("OPENAI_API_KEY encontrada")

	baseUrlOpenAi := "https://api.openai.com/v1"
	envBaseURL := os.Getenv("OPENAI_BASE_URL")
	if envBaseURL != "" {
		baseUrlOpenAi = envBaseURL
		log.Println("Usando OPENAI_BASE_URL fornecida:", baseUrlOpenAi)
	} else {
		log.Println("OPENAI_BASE_URL não encontrada, usando URL padrão:", baseUrlOpenAi)
	}

	log.Println("Iniciando serviço de transcrição OpenAI...")
	transcriptionService := openai.NewAudioTranscriptionService(openAIOption.WithAPIKey(
		apiKeyOpenAI,
	),
		openAIOption.WithBaseURL(
			baseUrlOpenAi,
		))
	log.Println("Serviço de transcrição OpenAI inicializado")

	apiKeyGemini := os.Getenv("GEMINI_API_KEY")
	if apiKeyGemini == "" {
		log.Fatal("GEMINI_API_KEY environment variable is not set")
	}
	log.Println("GEMINI_API_KEY encontrada")

	log.Println("Iniciando cliente Gemini...")
	sumupService, err := genai.NewClient(context.Background(), geminiOption.WithAPIKey(apiKeyGemini))
	if err != nil {
		log.Fatalf("Erro ao criar cliente Gemini: %v", err)
	}
	log.Println("Cliente Gemini criado com sucesso")

	registerDependencies(api, transcriptionService, sumupService)
}

func registerDaos(transcriptionService openai.AudioTranscriptionService, sumupService *genai.Client) {
	log.Println("Registrando DAOs...")
	transcribingDao = dao.NewTranscribingDao(transcriptionService)
	sumupDao = dao.NewSumupDao(sumupService)
	log.Println("DAOs registrados com sucesso")
}
func registerServices() {
	log.Println("Registrando serviços...")
	transcribingService = service.NewTranscribingService(transcribingDao)
	sumupService = service.NewSumupService(sumupDao, transcribingService)
	log.Println("Serviços registrados com sucesso")
}
func registerControllers(api *echo.Group) {
	log.Println("Registrando controllers...")
	transcribingController = controller.NewTranscribingController(api, transcribingService)
	sumupController = controller.NewSumupController(api, sumupService)
	log.Println("Controllers registrados com sucesso")
}
