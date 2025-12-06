package dao

import (
	"mime/multipart"

	"github.com/lrstanley/go-ytdlp"
)

type DownloadDao struct {
	downloadClient *ytdlp.Command	
}

func NewDownloadDao() *DownloadDao {
	return &DownloadDao{}
}

func (d *DownloadDao) DownloadVideoAudioFromUrl(url string) (*multipart.FileHeader, error) {
	return &multipart.FileHeader{}, nil
}