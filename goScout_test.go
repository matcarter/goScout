package goScout

import (
	"github.com/matcarter/goScout/configs"
	"github.com/stretchr/testify/require"
	"testing"
)

var configFile = "configs/config.json"

func TestNewScout(t *testing.T) {
	config, err := configs.GetConfig(configFile)
	require.Nil(t, err)

	scout := NewScout(config.APIKey, config.Region)
	require.NotNil(t, scout)
}
