package v1serializers

import (
	"time"

	"github.com/hoangmirs/go-scraper/models"
)

type KeywordList struct {
	Keywords []*models.Keyword
}

type KeywordResponse struct {
	Id        uint   `jsonapi:"primary,keyword"`
	Keyword   string `jsonapi:"attr,keyword"`
	Status    string `jsonapi:"attr,status"`
	CreatedAt string `jsonapi:"attr,created_at"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

func (serializer *KeywordList) Data() []*KeywordResponse {
	var data []*KeywordResponse

	for _, keyword := range serializer.Keywords {
		data = append(data, createKeywordResponse(keyword))
	}

	return data
}

func createKeywordResponse(keyword *models.Keyword) *KeywordResponse {
	return &KeywordResponse{
		Id:        keyword.Id,
		Keyword:   keyword.Keyword,
		Status:    string(keyword.Status),
		CreatedAt: keyword.CreatedAt.Format(time.RFC3339),
		UpdatedAt: keyword.UpdatedAt.Format(time.RFC3339),
	}
}