package scraper_test

import (
	"github.com/hoangmirs/go-scraper/models"
	"github.com/hoangmirs/go-scraper/services/scraper"
	"github.com/hoangmirs/go-scraper/tests/fabricators"
	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	"github.com/beego/beego/v2/client/orm"
	"github.com/bxcodec/faker/v3"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/gocolly/colly/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ScraperService", func() {
	AfterEach(func() {
		TruncateTables("keyword_result")
	})

	Describe("#Run", func() {
		Context("given valid attributes", func() {
			It("gets the result", func() {
				cassetteName := "scraper/success_valid_attributes"
				rec, err := recorder.New(CassettePath(cassetteName))
				if err != nil {
					Fail(err.Error())
				}

				defer func() {
					err := rec.Stop()
					if err != nil {
						Fail(err.Error())
					}
				}()

				user := fabricators.FabricateUser(faker.Email(), faker.Password())
				keyword := &models.Keyword{
					Keyword: "iphone 12",
					User:    user,
				}
				collector := colly.NewCollector()
				collector.WithTransport(rec)
				service := scraper.ScraperService{Keyword: keyword, Collector: collector}
				err = service.Run()
				if err != nil {
					Fail(err.Error())
				}

				Expect(len(service.GetParsingResult().NonAdwordLinks)).To(BeNumerically(">", 0))
				Expect(len(service.GetParsingResult().AdwordLinks)).To(BeNumerically(">", 0))
				Expect(len(service.GetParsingResult().ShopAdwordLinks)).To(BeNumerically(">", 0))
				Expect(service.GetParsingResult().HTMLCode).NotTo(BeNil())
			})

			It("save the result to DB", func() {
				cassetteName := "scraper/success_valid_attributes"
				rec, err := recorder.New(CassettePath(cassetteName))
				if err != nil {
					Fail(err.Error())
				}

				defer func() {
					err := rec.Stop()
					if err != nil {
						Fail(err.Error())
					}
				}()

				searchKeyword := "iphone 12"
				user := fabricators.FabricateUser(faker.Email(), faker.Password())
				keyword := &models.Keyword{
					Keyword: searchKeyword,
					User:    user,
				}
				collector := colly.NewCollector()
				collector.WithTransport(rec)
				service := scraper.ScraperService{Keyword: keyword, Collector: collector}
				err = service.Run()
				if err != nil {
					Fail(err.Error())
				}

				savedKeyword := models.Keyword{}

				o := orm.NewOrm()
				err = o.QueryTable("keyword").Filter("keyword", searchKeyword).One(&savedKeyword)
				if err != nil {
					Fail(err.Error())
				}

				Expect(savedKeyword.Id).To(BeNumerically(">", 0))
				Expect(savedKeyword.Keyword).To(Equal(searchKeyword))
			})
		})
	})

	Context("given invalid attributes", func() {
		Context("given a blank keyword", func() {
			It("returns an error", func() {
				user := fabricators.FabricateUser(faker.Email(), faker.Password())
				keyword := &models.Keyword{
					Keyword: "",
					User:    user,
				}
				service := scraper.ScraperService{Keyword: keyword}
				err := service.Run()

				Expect(err.Error()).To(Equal("Keyword required"))
			})
		})

		Context("given NO user object", func() {
			It("returns an error", func() {
				keyword := &models.Keyword{
					Keyword: "iphone 12",
				}
				service := scraper.ScraperService{Keyword: keyword}
				err := service.Run()

				Expect(err.Error()).To(Equal("User required"))
			})
		})
	})
})
