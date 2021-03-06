package v1serializers_test

import (
	"time"

	"github.com/hoangmirs/go-scraper/models"
	v1serializers "github.com/hoangmirs/go-scraper/serializers/v1"
	"github.com/hoangmirs/go-scraper/tests/fabricators"
	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeywordListSerializer", func() {
	AfterEach(func() {
		TruncateTables("user", "keyword")
	})

	Describe("Data", func() {
		Context("given a valid attributes", func() {
			It("returns correct data", func() {
				user := fabricators.FabricateUser(faker.Email(), faker.Password())
				keyword1 := fabricators.FabricateKeyword(faker.Word(), user)
				keyword2 := fabricators.FabricateKeyword(faker.Word(), user)
				query := map[string]interface{}{"user_id": user.Id}

				keywords, err := models.GetKeywords(query)
				if err != nil {
					Fail("Error when getting keywords: %v" + err.Error())
				}

				serializer := v1serializers.KeywordList{
					Keywords: keywords,
				}

				data := serializer.Data()

				Expect(data[0].Id).To(Equal(keyword1.Id))
				Expect(data[0].Keyword).To(Equal(keyword1.Keyword))
				Expect(data[0].Status).To(Equal(string(keyword1.Status)))
				Expect(data[0].CreatedAt).To(Equal(keyword1.CreatedAt.UTC().Format(time.RFC3339)))
				Expect(data[0].UpdatedAt).To(Equal(keyword1.UpdatedAt.UTC().Format(time.RFC3339)))
				Expect(data[1].Id).To(Equal(keyword2.Id))
				Expect(data[1].Keyword).To(Equal(keyword2.Keyword))
				Expect(data[1].Status).To(Equal(string(keyword2.Status)))
				Expect(data[1].CreatedAt).To(Equal(keyword2.CreatedAt.UTC().Format(time.RFC3339)))
				Expect(data[1].UpdatedAt).To(Equal(keyword2.UpdatedAt.UTC().Format(time.RFC3339)))
			})
		})
	})
})
