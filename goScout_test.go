package goScout

import (
	"github.com/matcarter/goScout/internal/configuration"
	"github.com/stretchr/testify/require"
	"testing"
)

var configFile = "internal/configuration/init.json"

func TestNewScout(t *testing.T) {
	config, err := configuration.GetConfig(configFile)
	require.Nil(t, err)

	scout, err := NewScout(config.APIKey, config.Region)
	require.Nil(t, err)
	require.NotNil(t, scout)
}
