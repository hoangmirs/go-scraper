package presenters_test

import (
	"encoding/json"

	"github.com/hoangmirs/go-scraper/models"
	"github.com/hoangmirs/go-scraper/presenters"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeywordPresenter", func() {
	Describe("#KeywordLinks", func() {
		It("returns correct keyword links", func() {
			nonAdwordLinks, _ := json.Marshal([]string{"non-adword-link"})
			adwordLinks, _ := json.Marshal([]string{"adword-link"})
			shopAdwordLinks, _ := json.Marshal([]string{"shop-adword-link"})

			keyword := &models.Keyword{
				Keyword:              "keyword",
				Status:               models.Pending,
				LinksCount:           3,
				NonAdwordLinks:       string(nonAdwordLinks),
				NonAdwordLinksCount:  1,
				AdwordLinks:          string(adwordLinks),
				AdwordLinksCount:     1,
				ShopAdwordLinks:      string(shopAdwordLinks),
				ShopAdwordLinksCount: 1,
				HtmlCode:             "HTML",
			}

			keywordPresenter := presenters.KeywordPresenter{Keyword: keyword}
			keywordPresenter.ConvertKeywordLinks()
			links := keywordPresenter.Links

			Expect(links.NonAdwordLinks[0]).To(Equal("non-adword-link"))
			Expect(links.AdwordLinks[0]).To(Equal("adword-link"))
			Expect(links.ShopAdwordLinks[0]).To(Equal("shop-adword-link"))
		})
	})
})
