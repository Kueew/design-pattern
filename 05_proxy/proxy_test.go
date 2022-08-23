package proxy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserProxyLogin(t *testing.T) {
	proxy := NewUserProxy(&User{})
	err := proxy.Login("user1", "password")
	require.Nil(t, err)
}
