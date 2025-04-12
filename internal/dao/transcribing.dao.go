package dao

import (
	"context"
	"io"

	"github.com/8bury/sumup4me/internal/model"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/packages/param"
)

type TranscribingDao struct {
	cliente openai.AudioTranscriptionService
}

func NewTranscribingDao(cliente openai.AudioTranscriptionService) *TranscribingDao {
	return &TranscribingDao{
		cliente: cliente,
	}
}

func (d *TranscribingDao) TranscribeAudio(file io.Reader) (*model.Transcription, error) {
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
	transcription, err := d.cliente.New(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return &model.Transcription{
		Transcription: transcription.Text,
	}, nil
}
