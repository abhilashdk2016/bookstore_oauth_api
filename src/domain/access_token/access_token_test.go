package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstant(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "Expiration time must be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken(0)
	assert.False(t, at.IsExpired(), "Brand new Access Token Should Not be Expired")

	assert.Equal(t, "", at.AccessToken, "New Access Token should not have defined access token id")

	assert.True(t, at.UserId == 0, "New Access Token should not have an associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "Empty access token should be expired by default")
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "Access token expiring 3 hours from now should not be expired")
}
