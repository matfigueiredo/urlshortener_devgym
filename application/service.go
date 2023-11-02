package service

import (
	"crypto/sha256"
	"encoding/base64"

	"github.com/matfigueiredo/urlshortener_devgym/domain"
)

type URLService struct {
	repo domain.URLRepository
}

func NewURLService(r domain.URLRepository) *URLService {
	return &URLService{repo: r}
}

func (s *URLService) ShortenURL(original string) (domain.URL, error) {
	code, err := generateShortCode(original)
	if err != nil {
		return domain.URL{}, err
	}

	url := domain.URL{
		Original: original,
		Code:     code,
	}
	err = s.repo.Save(url)
	if err != nil {
		return domain.URL{}, err
	}

	return url, nil
}

func (s *URLService) GetOriginalURL(code string) (domain.URL, error) {
	return s.repo.FindByCode(code)
}

func generateShortCode(original string) (string, error) {
	hasher := sha256.New()
	hasher.Write([]byte(original))
	encoded := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return encoded[:6], nil
}
