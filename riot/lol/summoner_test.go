package lol

import (
	"github.com/KnutZuidema/golio/api"
	"github.com/matcarter/goScout/configs"
	"github.com/matcarter/goScout/internal"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
)

var configFile = "../../configs/config.json"

func TestNewClient(t *testing.T) {
	config, err := configs.GetConfig(configFile)
	require.Nil(t, err)
	require.NotNil(t, config)

	base := internal.NewClient(config.APIKey, api.Region(config.Region), logrus.StandardLogger())
	require.NotNil(t, base)

	client := NewClient(base)
	require.NotNil(t, client)
}

func TestSummonerClient_ScoutByName(t *testing.T) {
	config, err := configs.GetConfig(configFile)
	require.Nil(t, err)
	require.NotNil(t, config)

	base := internal.NewClient(config.APIKey, api.Region(config.Region), logrus.StandardLogger())
	require.NotNil(t, base)

	lolClient := NewClient(base)
	require.NotNil(t, lolClient)

	report, err := lolClient.Summoner.ScoutByName("Bully")
	require.Nil(t, err)
	require.NotNil(t, report)
}

func TestSummonerClient_AnalyzeData(t *testing.T) {

}
