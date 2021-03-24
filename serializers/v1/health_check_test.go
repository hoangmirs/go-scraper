package v1serializers_test

import (
	v1serializers "github.com/hoangmirs/go-scraper/serializers/v1"
	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HealthCheckSerializer", func() {
	AfterEach(func() {
		TruncateTables("user", "keyword")
	})

	Describe("Data", func() {
		Context("given a valid data", func() {
			It("returns correct data", func() {
				serializer := v1serializers.HealthCheck{
					HealthCheck: true,
				}

				data := serializer.Data()

				Expect(data.Id).To(Equal(0))
				Expect(data.Success).To(Equal(true))
			})
		})
	})
})
