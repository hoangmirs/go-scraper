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

	Describe("#UpdateKeyword", func() {
		Context("given valid attributes", func() {
			It("updates the record successfully", func() {
				user := fabricators.FabricateUser(faker.Email(), faker.Password())
				keyword := fabricators.FabricateKeyword("iphone 12", user)
				keyword.Keyword = "macbook pro"

				err := models.UpdateKeyword(keyword)
				if err != nil {
					Fail("Failed to update keyword: " + err.Error())
				}

				updatedKeyword, err := models.GetKeywordByID(int64(keyword.Id))
				if err != nil {
					Fail("Failed to get keyword: " + err.Error())
				}

				Expect(updatedKeyword.Keyword).To(Equal("macbook pro"))
			})

			It("does NOT return error", func() {
				user := fabricators.FabricateUser(faker.Email(), faker.Password())
				keyword := fabricators.FabricateKeyword("iphone 12", user)
				keyword.Keyword = "macbook pro"

				err := models.UpdateKeyword(keyword)

				Expect(err).To(BeNil())
			})
		})

		Context("given invalid attributes", func() {
			It("returns an error", func() {
				user := fabricators.FabricateUser(faker.Email(), faker.Password())
				keyword := fabricators.FabricateKeyword("iphone 12", user)
				keyword.NonAdwordLinks = "invalid links"

				err := models.UpdateKeyword(keyword)

				Expect(err.Error()).To(Equal("pq: invalid input syntax for type json"))
			})
		})
	})

	Describe("#GetKeywords", func() {
		Context("given a valid query", func() {
			It("returns uploaded keywords of user with the correct ordering", func() {
				user := fabricators.FabricateUser(faker.Email(), faker.Password())
				keyword1 := fabricators.FabricateKeyword("keyword 1", user)
				keyword2 := fabricators.FabricateKeyword("keyword 2", user)

				query := map[string]interface{}{
					"user_id": user.Id,
					"order":   "-id",
					"limit":   2,
					"offset":  0,
				}

				keywords, err := models.GetKeywords(query)
				if err != nil {
					Fail("Failed to get keywords: " + err.Error())
				}

				Expect(keywords[0].Id).To(Equal(keyword2.Id))
				Expect(keywords[1].Id).To(Equal(keyword1.Id))
			})
		})

		Context("given an invalid query", func() {
			It("panics the application", func() {
				query := map[string]interface{}{
					"invalid": "invalid",
				}

				Expect(func() {
					_, _ = models.GetKeywords(query)
				}).To(Panic())
			})
		})
	})

	Describe("#GetKeywordsCount", func() {
		Context("given a valid query", func() {
			Context("given a correct user ID", func() {
				It("returns the number of keywords of the user", func() {
					user := fabricators.FabricateUser(faker.Email(), faker.Password())
					_ = fabricators.FabricateKeyword("keyword 1", user)

					query := map[string]interface{}{
						"user_id": user.Id,
					}

					keywordCount, err := models.GetKeywordsCount(query)
					if err != nil {
						Fail("Failed to get keyword count: " + err.Error())
					}

					Expect(keywordCount).To(Equal(int64(1)))
				})
			})

			Context("given an incorrect user ID", func() {
				It("returns 0", func() {
					user := fabricators.FabricateUser(faker.Email(), faker.Password())
					_ = fabricators.FabricateKeyword("keyword 1", user)

					query := map[string]interface{}{
						"user_id": 0,
					}

					keywordCount, err := models.GetKeywordsCount(query)
					if err != nil {
						Fail("Failed to get keyword count: " + err.Error())
					}

					Expect(keywordCount).To(Equal(int64(0)))
				})
			})
		})

		Context("given an invalid query", func() {
			It("panics the application", func() {
				query := map[string]interface{}{
					"invalid": "invalid",
				}

				Expect(func() {
					_, _ = models.GetKeywordsCount(query)
				}).To(Panic())
			})
		})
	})

	Describe("#GetKeywordByQuery", func() {
		Context("given a valid query", func() {
			Context("given correct keyword and user ID", func() {
				It("returns correct keyword", func() {
					user := fabricators.FabricateUser(faker.Email(), faker.Password())
					savedKeyword := fabricators.FabricateKeyword("keyword", user)

					query := map[string]interface{}{
						"keyword": savedKeyword.Keyword,
						"user_id": user.Id,
					}
					keyword, err := models.GetKeywordByQuery(query)
					if err != nil {
						Fail("Failed to get keyword: " + err.Error())
					}

					Expect(keyword.Id).To(Equal(savedKeyword.Id))
				})

				It("does NOT return error", func() {
					user := fabricators.FabricateUser(faker.Email(), faker.Password())
					savedKeyword := fabricators.FabricateKeyword("keyword", user)

					query := map[string]interface{}{
						"keyword": savedKeyword.Keyword,
						"user_id": user.Id,
					}
					_, err := models.GetKeywordByQuery(query)

					Expect(err).To(BeNil())
				})
			})

			Context("given a wrong keyword", func() {
				It("does NOT return keyword", func() {
					user := fabricators.FabricateUser(faker.Email(), faker.Password())
					_ = fabricators.FabricateKeyword("keyword", user)

					query := map[string]interface{}{
						"keyword": "wrong keyword",
					}
					keyword, _ := models.GetKeywordByQuery(query)

					Expect(keyword).To(BeNil())
				})

				It("returns an error", func() {
					user := fabricators.FabricateUser(faker.Email(), faker.Password())
					_ = fabricators.FabricateKeyword("keyword", user)

					query := map[string]interface{}{
						"keyword": "wrong keyword",
					}
					_, err := models.GetKeywordByQuery(query)

					Expect(err.Error()).To(Equal("Keyword not found"))
				})
			})
		})

		Context("given an invalid query", func() {
			It("panics the application", func() {
				query := map[string]interface{}{
					"invalid": "invalid",
				}

				Expect(func() {
					_, _ = models.GetKeywordByQuery(query)
				}).To(Panic())
			})
		})
	})
})
