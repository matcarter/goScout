package configuration

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetConfig(t *testing.T) {
	configFile := "init.json"
	config, err := GetConfig(configFile)
	require.Nil(t, err)
	require.NotNil(t, config)
}

func TestGetConfig_Read_Error(t *testing.T) {
	configFile := "definitely_not_a_config_file.json"
	config, err := GetConfig(configFile)
	require.NotNil(t, err)
	require.Nil(t, config)
}

func TestGetConfig_Unmarshal_Error(t *testing.T) {

}
