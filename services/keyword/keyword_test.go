package keyword_test

import (
	"github.com/hoangmirs/go-scraper/conf"
	"github.com/hoangmirs/go-scraper/services/keyword"
	. "github.com/hoangmirs/go-scraper/tests/custom_matchers"
	"github.com/hoangmirs/go-scraper/tests/fabricators"
	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeywordService", func() {
	AfterEach(func() {
		TruncateTables("user", "keyword")
		DeleteRedisJobs()
	})

	Describe("#Run", func() {
		Context("given valid attributes", func() {
			It("does NOT return error", func() {
				user := fabricators.FabricateUser(faker.Email(), faker.Password())

				service := keyword.KeywordService{
					Keywords: []string{"iphone 12"},
					User:     user,
				}

				err := service.Run()

				Expect(err).To(BeNil())
			})

			It("enqueues the job", func() {
				user := fabricators.FabricateUser(faker.Email(), faker.Password())

				service := keyword.KeywordService{
					Keywords: []string{"iphone 12"},
					User:     user,
				}

				err := service.Run()
				if err != nil {
					Fail(err.Error())
				}

				workerClient := GetWorkerClient()

				Expect(workerClient).To(EnqueueJob(conf.GetString("scraperJobName"), 1))
				Expect(err).To(BeNil())
			})
		})

		Context("given invalid attributes", func() {
			Context("given empty keywords", func() {
				It("returns error", func() {
					user := fabricators.FabricateUser(faker.Email(), faker.Password())

					service := keyword.KeywordService{
						Keywords: []string{},
						User:     user,
					}

					err := service.Run()

					Expect(err.Error()).To(Equal("Keywords are empty"))
				})
			})

			Context("given empty keywords", func() {
				It("returns error", func() {
					service := keyword.KeywordService{
						Keywords: []string{"iphone 12"},
						User:     nil,
					}

					err := service.Run()

					Expect(err.Error()).To(Equal("User object required"))
				})
			})
		})
	})
})
