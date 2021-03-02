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
		TruncateTables("user", "keyword")
	})

	Describe("#Run", func() {
		Context("given valid attributes", func() {
			It("returns the parsing result", func() {
				// TODO : Refactor VCR recorder and put it into test helper
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

				user, err := fabricators.FabricateUser(faker.Email(), faker.Password())
				if err != nil {
					Fail(err.Error())
				}

				keyword, err := fabricators.FabricateKeyword("iphone 12", user)
				if err != nil {
					Fail(err.Error())
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

			It("saves the result to DB", func() {
				// TODO : Refactor VCR recorder and put it into test helper
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
				user, err := fabricators.FabricateUser(faker.Email(), faker.Password())
				if err != nil {
					Fail(err.Error())
				}

				keyword, err := fabricators.FabricateKeyword("iphone 12", user)
				if err != nil {
					Fail(err.Error())
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
				Expect(savedKeyword.Status).To(Equal(models.Processed))
				Expect(len(savedKeyword.HtmlCode)).To(BeNumerically(">", 0))
				Expect(savedKeyword.LinksCount).To(BeNumerically(">", 0))
			})
		})
	})

	Context("given invalid attributes", func() {
		Context("given no keyword object", func() {
			It("returns an error", func() {
				service := scraper.ScraperService{Keyword: nil}
				err := service.Run()

				Expect(err.Error()).To(Equal("Keyword object required"))
			})
		})
	})
})
