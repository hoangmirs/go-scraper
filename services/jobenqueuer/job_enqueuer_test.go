package jobenqueuer_test

import (
	"github.com/hoangmirs/go-scraper/conf"
	"github.com/hoangmirs/go-scraper/services/jobenqueuer"
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

	Describe("#EnqueueKeyword", func() {
		Context("given a valid keyword", func() {
			It("does NOT return error", func() {
				user := fabricators.FabricateUser(faker.Email(), faker.Password())
				keyword := fabricators.FabricateKeyword("iphone 12", user)

				err := jobenqueuer.EnqueueKeyword(keyword)

				Expect(err).To(BeNil())
			})

			It("enqueues the job", func() {
				user := fabricators.FabricateUser(faker.Email(), faker.Password())
				keyword := fabricators.FabricateKeyword("iphone 12", user)

				err := jobenqueuer.EnqueueKeyword(keyword)
				if err != nil {
					Fail(err.Error())
				}

				workerClient := GetWorkerClient()

				Expect(workerClient).To(EnqueueJob(conf.GetString("scraperJobName"), 1))
				Expect(err).To(BeNil())
			})
		})

		Context("given a nil keyword", func() {
			It("returns an error", func() {
				err := jobenqueuer.EnqueueKeyword(nil)

				Expect(err.Error()).To(Equal("Keyword cannot be nil"))
			})
		})
	})
})
