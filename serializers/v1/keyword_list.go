package v1serializers

import (
	"time"

	"github.com/hoangmirs/go-scraper/models"
)

type KeywordList struct {
	Keywords []*models.Keyword

	Pagination
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
		CreatedAt: keyword.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt: keyword.UpdatedAt.UTC().Format(time.RFC3339),
	}
}
