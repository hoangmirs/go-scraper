package presenters

import (
	"encoding/json"

	"github.com/hoangmirs/go-scraper/models"

	"github.com/beego/beego/v2/core/logs"
)

type KeywordPresenter struct {
	Keyword *models.Keyword
	Links   *KeywordLinks
}

type KeywordLinks struct {
	NonAdwordLinks  []string
	AdwordLinks     []string
	ShopAdwordLinks []string
}

func (kp *KeywordPresenter) ConvertKeywordLinks() {
	var nonAdwordLinks []string
	if kp.Keyword.NonAdwordLinks != "" {
		err := json.Unmarshal([]byte(kp.Keyword.NonAdwordLinks), &nonAdwordLinks)
		if err != nil {
			logs.Error("Cannot unmarshal NonAdwordLinks", err)
		}
	}

	var adwordLinks []string
	if kp.Keyword.AdwordLinks != "" {
		err := json.Unmarshal([]byte(kp.Keyword.AdwordLinks), &adwordLinks)
		if err != nil {
			logs.Error("Cannot unmarshal AdwordLinks", err)
		}
	}

	var shopAdwordLinks []string
	if kp.Keyword.ShopAdwordLinks != "" {
		err := json.Unmarshal([]byte(kp.Keyword.ShopAdwordLinks), &shopAdwordLinks)
		if err != nil {
			logs.Error("Cannot unmarshal ShopAdwordLinks", err)
		}
	}

	kp.Links = &KeywordLinks{
		NonAdwordLinks:  nonAdwordLinks,
		AdwordLinks:     adwordLinks,
		ShopAdwordLinks: shopAdwordLinks,
	}
}
