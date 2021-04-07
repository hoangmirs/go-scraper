package v1serializers_test

import (
	v1serializers "github.com/hoangmirs/go-scraper/serializers/v1"
	"github.com/hoangmirs/go-scraper/tests/fabricators"
	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	"github.com/google/uuid"
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
				oauthClient := fabricators.FabricateOAuthClient(uuid.New().String(), uuid.New().String())

				serializer := v1serializers.OAuthClient{
					OAuthClient: oauthClient,
				}

				data := serializer.Data()

				Expect(data.ID).To(Equal(oauthClient.ID))
				Expect(data.ClientID).To(Equal(oauthClient.ID))
				Expect(data.ClientSecret).To(Equal(oauthClient.Secret))
			})
		})
	})
})
