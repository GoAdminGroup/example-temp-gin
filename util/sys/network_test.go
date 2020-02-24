package sys

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNetworkLocalIP(t *testing.T) {
	// mock
	// do
	ipv4, err := NetworkLocalIP()
	if err != nil {
		t.Fatalf("TestNetworkLocalIP test error: %v", err)
	}
	// verify
	t.Logf("NetworkLocalIP get: %v", ipv4)
	assert.NotNil(t, ipv4)
}
