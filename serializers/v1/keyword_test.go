package v1serializers_test

import (
	"time"

	"github.com/hoangmirs/go-scraper/presenters"
	v1serializers "github.com/hoangmirs/go-scraper/serializers/v1"
	"github.com/hoangmirs/go-scraper/tests/fabricators"
	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeywordSerializer", func() {
	AfterEach(func() {
		TruncateTables("user", "keyword")
	})

	Describe("Data", func() {
		Context("given a valid attributes", func() {
			It("returns correct data", func() {
				user := fabricators.FabricateUser(faker.Email(), faker.Password())
				keyword := fabricators.FabricateKeyword(faker.Word(), user)
				keywordPresenter := presenters.KeywordPresenter{Keyword: keyword}
				keywordPresenter.ConvertKeywordLinks()

				serializer := v1serializers.Keyword{
					Keyword: keyword,
					Links:   keywordPresenter.Links,
				}

				data := serializer.Data()

				Expect(data.Id).To(Equal(keyword.Id))
				Expect(data.Keyword).To(Equal(keyword.Keyword))
				Expect(data.Status).To(Equal(string(keyword.Status)))
				Expect(data.CreatedAt).To(Equal(keyword.CreatedAt.UTC().Format(time.RFC3339)))
				Expect(data.UpdatedAt).To(Equal(keyword.UpdatedAt.UTC().Format(time.RFC3339)))
				Expect(data.LinkCount).To(Equal(0))
				Expect(data.NonAdwordLinksCount).To(Equal(0))
				Expect(data.AdwordLinksCount).To(Equal(0))
				Expect(data.ShopAdwordLinksCount).To(Equal(0))
			})
		})
	})
})
