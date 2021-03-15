package presenters

import (
	"encoding/json"

	"github.com/hoangmirs/go-scraper/models"

	"github.com/beego/beego/v2/core/logs"
)

type KeywordPresenter struct {
	Keyword *models.Keyword
}

type KeywordLinks struct {
	NonAdwordLinks  []string
	AdwordLinks     []string
	ShopAdwordLinks []string
}

func (kp *KeywordPresenter) KeywordLinks() KeywordLinks {
	var nonAdwordLinks []string
	err := json.Unmarshal([]byte(kp.Keyword.NonAdwordLinks), &nonAdwordLinks)
	if err != nil {
		logs.Error("Cannot unmarshal NonAdwordLinks", err)
	}

	var adwordLinks []string
	err = json.Unmarshal([]byte(kp.Keyword.AdwordLinks), &adwordLinks)
	if err != nil {
		logs.Error("Cannot unmarshal AdwordLinks", err)
	}

	var shopAdwordLinks []string
	err = json.Unmarshal([]byte(kp.Keyword.ShopAdwordLinks), &shopAdwordLinks)
	if err != nil {
		logs.Error("Cannot unmarshal ShopAdwordLinks", err)
	}

	return KeywordLinks{
		NonAdwordLinks:  nonAdwordLinks,
		AdwordLinks:     adwordLinks,
		ShopAdwordLinks: shopAdwordLinks,
	}
}
