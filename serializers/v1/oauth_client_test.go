package v1serializers_test

import (
	v1serializers "github.com/hoangmirs/go-scraper/serializers/v1"
	oauthservice "github.com/hoangmirs/go-scraper/services/oauth"
	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OAuthClientSerializer", func() {
	AfterEach(func() {
		TruncateTables("oauth2_clients")
	})

	Describe("Data", func() {
		Context("given a valid OAuthClient", func() {
			It("returns correct data", func() {
				oauthClient, err := oauthservice.GenerateClient()
				if err != nil {
					Fail("Error when generating OAuth client: %v" + err.Error())
				}

				serializer := v1serializers.OAuthClient{
					OAuthClient: oauthClient,
				}

				data := serializer.Data()

				Expect(data.ClientID).To(Equal(oauthClient.ID))
				Expect(data.ClientSecret).To(Equal(oauthClient.Secret))
			})
		})
	})
})
