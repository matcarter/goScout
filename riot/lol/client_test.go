package lol

import (
	"testing"

	"github.com/KnutZuidema/golio/api"
	"github.com/matcarter/goScout/internal"
	"github.com/matcarter/goScout/internal/configuration"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	configFile := "../../configuration/init.json"
	config, err := configuration.GetConfig(configFile)
	require.Nil(t, err)
	require.NotNil(t, config)

	base, err := internal.NewClient(config.APIKey, api.Region(config.Region), logrus.StandardLogger())
	require.NotNil(t, err)
	require.NotNil(t, base)

	client, err := NewClient(base)
	require.NotNil(t, err)
	require.NotNil(t, client)
}
