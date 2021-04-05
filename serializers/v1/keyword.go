package v1serializers

import (
	"time"

	"github.com/hoangmirs/go-scraper/models"
	"github.com/hoangmirs/go-scraper/presenters"
)

type Keyword struct {
	Keyword *models.Keyword
	Links   *presenters.KeywordLinks
}

type KeywordResponse struct {
	Id                   uint     `jsonapi:"primary,keyword"`
	Keyword              string   `jsonapi:"attr,keyword"`
	Status               string   `jsonapi:"attr,status"`
	CreatedAt            string   `jsonapi:"attr,created_at"`
	UpdatedAt            string   `jsonapi:"attr,updated_at"`
	LinkCount            int      `jsonapi:"attr,link_count,omitempty"`
	NonAdwordLinksCount  int      `jsonapi:"attr,non_adword_link_count,omitempty"`
	NonAdwordLinks       []string `jsonapi:"attr,non_adword_links,omitempty"`
	AdwordLinksCount     int      `jsonapi:"attr,adword_link_count,omitempty"`
	AdwordLinks          []string `jsonapi:"attr,adword_links,omitempty"`
	ShopAdwordLinksCount int      `jsonapi:"attr,shop_adword_link_count,omitempty"`
	ShopAdwordLinks      []string `jsonapi:"attr,shop_adword_links,omitempty"`
	HTMLCode             string   `jsonapi:"attr,html_code,omitempty"`
}

func (serializer *Keyword) Data() *KeywordResponse {
	return &KeywordResponse{
		Id:                   serializer.Keyword.Id,
		Keyword:              serializer.Keyword.Keyword,
		Status:               string(serializer.Keyword.Status),
		CreatedAt:            serializer.Keyword.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt:            serializer.Keyword.UpdatedAt.UTC().Format(time.RFC3339),
		LinkCount:            serializer.Keyword.LinksCount,
		NonAdwordLinksCount:  serializer.Keyword.NonAdwordLinksCount,
		NonAdwordLinks:       serializer.Links.NonAdwordLinks,
		AdwordLinksCount:     serializer.Keyword.AdwordLinksCount,
		AdwordLinks:          serializer.Links.AdwordLinks,
		ShopAdwordLinksCount: serializer.Keyword.ShopAdwordLinksCount,
		ShopAdwordLinks:      serializer.Links.ShopAdwordLinks,
		HTMLCode:             serializer.Keyword.HtmlCode,
	}
}
