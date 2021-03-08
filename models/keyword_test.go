package models_test

import (
	"github.com/hoangmirs/go-scraper/models"
	"github.com/hoangmirs/go-scraper/tests/fabricators"
	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	AfterEach(func() {
		TruncateTables("keyword")
	})

	Describe("#CreateKeyword", func() {
		Context("given valid attributes", func() {
			It("returns the keyword ID", func() {
				user := fabricators.FabricateUser(faker.Email(), faker.Password())

				keyword := models.Keyword{
					Keyword:              "keyword",
					LinksCount:           4,
					NonAdwordLinks:       `["link1", "link2"]`,
					NonAdwordLinksCount:  2,
					AdwordLinks:          `["link"]`,
					AdwordLinksCount:     1,
					ShopAdwordLinks:      `["link"]`,
					ShopAdwordLinksCount: 1,
					HtmlCode:             "<html></html>",
					User:                 user,
				}

				keywordID, err := models.CreateKeyword(&keyword)
				if err != nil {
					Fail("Failed to create keyword: " + err.Error())
				}

				Expect(keywordID).To(BeNumerically(">", 0))
			})

			It("does NOT return error", func() {
				user := fabricators.FabricateUser(faker.Email(), faker.Password())

				keyword := models.Keyword{
					Keyword:              "keyword",
					LinksCount:           4,
					NonAdwordLinks:       `["link1", "link2"]`,
					NonAdwordLinksCount:  2,
					AdwordLinks:          `["link"]`,
					AdwordLinksCount:     1,
					ShopAdwordLinks:      `["link"]`,
					ShopAdwordLinksCount: 1,
					HtmlCode:             "<html></html>",
					User:                 user,
				}

				_, err := models.CreateKeyword(&keyword)
				if err != nil {
					Fail("Failed to create keyword: " + err.Error())
				}

				Expect(err).To(BeNil())
			})
		})

		Context("given invalid attributes", func() {
			Context("given NO keyword attribute", func() {
				It("returns an error", func() {
					user := fabricators.FabricateUser(faker.Email(), faker.Password())

					keyword := models.Keyword{
						LinksCount:           4,
						NonAdwordLinks:       `["link1", "link2"]`,
						NonAdwordLinksCount:  2,
						AdwordLinks:          `["link"]`,
						AdwordLinksCount:     1,
						ShopAdwordLinks:      `["link"]`,
						ShopAdwordLinksCount: 1,
						HtmlCode:             "<html></html>",
						User:                 user,
					}

					_, err := models.CreateKeyword(&keyword)

					Expect(err.Error()).To(Equal("Keyword required"))
				})
			})

			Context("given NO user attribute", func() {
				It("returns an error", func() {
					keyword := models.Keyword{
						Keyword:              "keyword",
						LinksCount:           4,
						NonAdwordLinks:       `["link1", "link2"]`,
						NonAdwordLinksCount:  2,
						AdwordLinks:          `["link"]`,
						AdwordLinksCount:     1,
						ShopAdwordLinks:      `["link"]`,
						ShopAdwordLinksCount: 1,
						HtmlCode:             "<html></html>",
					}

					_, err := models.CreateKeyword(&keyword)

					Expect(err.Error()).To(Equal("User required"))
				})
			})
		})
	})

})
