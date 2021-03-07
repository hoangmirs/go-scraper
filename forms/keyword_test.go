package forms_test

import (
	"github.com/bxcodec/faker/v3"
	"github.com/hoangmirs/go-scraper/forms"
	"github.com/hoangmirs/go-scraper/tests/fabricators"
	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeywordForm", func() {
	AfterEach(func() {
		TruncateTables("keyword", "user")
	})

	Describe("Save", func() {
		Context("given valid attributes", func() {
			It("does NOT return error", func() {
				file, fileHeader, err := GetMultipartAttributesFromFile("tests/fixtures/files/valid.csv", "text/csv")
				if err != nil {
					Fail(err.Error())
				}

				user, err := fabricators.FabricateUser(faker.Email(), faker.Password())
				if err != nil {
					Fail("Failed to fabricate user: " + err.Error())
				}

				keywordFrom := forms.KeywordForm{
					File:       file,
					FileHeader: fileHeader,
					User:       user,
				}

				err = keywordFrom.Save()

				Expect(err).To(BeNil())
			})
		})
	})

	Context("given invalid attributes", func() {
		Context("given NO file", func() {
			It("returns error", func() {
				user, err := fabricators.FabricateUser(faker.Email(), faker.Password())
				if err != nil {
					Fail("Failed to fabricate user: " + err.Error())
				}

				keywordFrom := forms.KeywordForm{
					File:       nil,
					FileHeader: nil,
					User:       user,
				}

				err = keywordFrom.Save()

				Expect(err.Error()).To(Equal("File can not be empty"))
			})
		})

		Context("given invalid file type", func() {
			It("returns error", func() {
				file, fileHeader, err := GetMultipartAttributesFromFile("tests/fixtures/files/text.txt", "text/plain")
				if err != nil {
					Fail(err.Error())
				}

				user, err := fabricators.FabricateUser(faker.Email(), faker.Password())
				if err != nil {
					Fail("Failed to fabricate user: " + err.Error())
				}

				keywordFrom := forms.KeywordForm{
					File:       file,
					FileHeader: fileHeader,
					User:       user,
				}

				err = keywordFrom.Save()

				Expect(err.Error()).To(Equal("File type is not supported"))
			})
		})

		Context("given invalid keyword length", func() {
			It("returns error", func() {
				file, fileHeader, err := GetMultipartAttributesFromFile("tests/fixtures/files/invalid.csv", "text/csv")
				if err != nil {
					Fail(err.Error())
				}

				user, err := fabricators.FabricateUser(faker.Email(), faker.Password())
				if err != nil {
					Fail("Failed to fabricate user: " + err.Error())
				}

				keywordFrom := forms.KeywordForm{
					File:       file,
					FileHeader: fileHeader,
					User:       user,
				}

				err = keywordFrom.Save()

				Expect(err.Error()).To(Equal("CSV file only accepts from 1 to 1000 keywords"))
			})
		})

		Context("given NO user object", func() {
			It("returns error", func() {
				file, fileHeader, err := GetMultipartAttributesFromFile("tests/fixtures/files/valid.csv", "text/csv")
				if err != nil {
					Fail(err.Error())
				}

				keywordFrom := forms.KeywordForm{
					File:       file,
					FileHeader: fileHeader,
					User:       nil,
				}

				err = keywordFrom.Save()

				Expect(err.Error()).To(Equal("User can not be empty"))
			})
		})
	})
})
