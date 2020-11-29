package goScout

import (
	"github.com/matcarter/goScout/configuration"
	"github.com/stretchr/testify/require"
	"testing"
)

var configFile = "configuration/init.json"

func TestNewScout(t *testing.T) {
	config, err := configuration.GetConfig(configFile)
	require.Nil(t, err)

	scout := NewScout(config.APIKey, config.Region)
	require.NotNil(t, scout)
}
