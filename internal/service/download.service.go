package service

import (
	"mime/multipart"

	"github.com/8bury/sumup4me/internal/dao"
)

type DownloadService struct {
	downloadDao *dao.DownloadDao
}

func NewDownloadService(downloadDao *dao.DownloadDao) *DownloadService {
	return &DownloadService{
		downloadDao: downloadDao,
	}
}

func (s *DownloadService) DownloadVideoAudioFromUrl(url string) (*multipart.FileHeader, error) {
	return s.downloadDao.DownloadVideoAudioFromUrl(url)
}
