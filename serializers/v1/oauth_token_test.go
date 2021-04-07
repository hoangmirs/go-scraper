package v1serializers_test

import (
	"github.com/hoangmirs/go-scraper/models"
	v1serializers "github.com/hoangmirs/go-scraper/serializers/v1"
	"github.com/hoangmirs/go-scraper/tests/fabricators"
	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OAuthTokenSerializer", func() {
	AfterEach(func() {
		TruncateTables("oauth2_tokens")
	})

	Describe("Data", func() {
		Context("given a valid OAuthToken", func() {
			It("returns correct data", func() {
				user := fabricators.FabricateUser(faker.Email(), faker.Password())
				token := fabricators.FabricateToken(user, nil)
				oAuthToken := &models.OAuthToken{
					AccessToken:  token.Access,
					RefreshToken: token.Refresh,
					TokenType:    "Bearer",
					ExpiresIn:    uint64(token.AccessExpiresIn),
				}

				serializer := v1serializers.OAuthToken{
					OAuthToken: oAuthToken,
				}

				data := serializer.Data()

				Expect(data.ID).To(Equal(oAuthToken.AccessToken))
				Expect(data.AccessToken).To(Equal(oAuthToken.AccessToken))
				Expect(data.RefreshToken).To(Equal(oAuthToken.RefreshToken))
				Expect(data.ExpiresIn).To(Equal(oAuthToken.ExpiresIn))
				Expect(data.TokenType).To(Equal(oAuthToken.TokenType))
			})
		})
	})
})
