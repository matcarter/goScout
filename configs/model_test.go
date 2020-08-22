package configs

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var configFile = "config.json"

func TestGetConfig(t *testing.T) {
	config, err := GetConfig(configFile)
	require.Nil(t, err)
	require.NotNil(t, config)
}
