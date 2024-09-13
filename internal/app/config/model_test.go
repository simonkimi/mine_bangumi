package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInitConfigStruct(t *testing.T) {
	var serverConfig ServerConfig
	initConfigStruct(&serverConfig)
	assert.Equal(t, "0.0.0.0", serverConfig.Ipv4Host)
	assert.Equal(t, 7962, serverConfig.Ipv4Port)
	assert.Equal(t, "[::1]", serverConfig.Ipv6Host)
	assert.Equal(t, 7962, serverConfig.Ipv6Port)

	var envServerConfig ServerConfig
	_ = os.Setenv("MBG_SERVER_IPV4_HOST", "127.0.0.1")
	_ = os.Setenv("MBG_SERVER_IPV4_PORT", "8080")
	_ = os.Setenv("MBG_SERVER_IPV6_HOST", "::1")
	_ = os.Setenv("MBG_SERVER_IPV6_PORT", "8080")
	initConfigStruct(&envServerConfig)
	assert.Equal(t, "127.0.0.1", envServerConfig.Ipv4Host)
	assert.Equal(t, 8080, envServerConfig.Ipv4Port)
	assert.Equal(t, "::1", envServerConfig.Ipv6Host)
	assert.Equal(t, 8080, envServerConfig.Ipv6Port)
}
