package dao

import (
	"context"
	"io"
	"log"

	"github.com/8bury/sumup4me/internal/model"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/packages/param"
)

type TranscribingDao struct {
	cliente openai.AudioTranscriptionService
}

func NewTranscribingDao(cliente openai.AudioTranscriptionService) *TranscribingDao {
	log.Println("Inicializando TranscribingDao")
	return &TranscribingDao{
		cliente: cliente,
	}
}

func (d *TranscribingDao) TranscribeAudio(file io.Reader) (*model.Transcription, error) {
	log.Println("Iniciando transcrição de áudio usando a API da OpenAI")

	request := openai.AudioTranscriptionNewParams{
		File:                   file,
		Model:                  "whisper-1",
		ResponseFormat:         openai.AudioResponseFormatJSON,
		Language:               param.NewOpt("pt"),
		Prompt:                 param.Opt[string]{},
		Temperature:            param.Opt[float64]{},
		Include:                []openai.TranscriptionInclude{},
		TimestampGranularities: []string{},
	}

	log.Println("Configuração de transcrição: model=whisper-1, language=pt, format=JSON")
	log.Println("Enviando requisição para a API da OpenAI...")

	transcription, err := d.cliente.New(context.Background(), request)
	if err != nil {
		log.Printf("Erro ao transcrever áudio com a API da OpenAI: %v", err)
		return nil, err
	}

	log.Println("Resposta recebida da API da OpenAI com sucesso")

	return &model.Transcription{
		Transcription: transcription.Text,
	}, nil
}
