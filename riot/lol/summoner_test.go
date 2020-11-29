package lol

import (
	"github.com/KnutZuidema/golio/api"
	"github.com/matcarter/goScout/internal"
	"github.com/matcarter/goScout/internal/configuration"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
)

var configFile = "../../internal/configuration/init.json"

func TestSummonerClient_ScoutByName(t *testing.T) {
	config, err := configuration.GetConfig(configFile)
	require.Nil(t, err)
	require.NotNil(t, config)

	base, err := internal.NewClient(config.APIKey, api.Region(config.Region), logrus.StandardLogger())
	require.NotNil(t, err)
	require.NotNil(t, base)

	lolClient, err := NewClient(base)
	require.NotNil(t, err)
	require.NotNil(t, lolClient)

	report, err := lolClient.Summoner.ScoutByName("Bully")
	require.Nil(t, err)
	require.NotNil(t, report)
}

func TestSummonerClient_GatherDataByName(t *testing.T) {
	config, err := configuration.GetConfig(configFile)
	require.Nil(t, err)
	require.NotNil(t, config)

	base, err := internal.NewClient(config.APIKey, api.Region(config.Region), logrus.StandardLogger())
	require.NotNil(t, err)
	require.NotNil(t, base)

	lolClient, err := NewClient(base)
	require.NotNil(t, err)
	require.NotNil(t, lolClient)

	data, err := lolClient.Summoner.GatherDataByName("Bully")
	require.Nil(t, err)
	require.NotNil(t, data)
}

func TestSummonerClient_GatherDataByName_Name_Error(t *testing.T) {
	config, err := configuration.GetConfig(configFile)
	require.Nil(t, err)
	require.NotNil(t, config)

	base, err := internal.NewClient(config.APIKey, api.Region(config.Region), logrus.StandardLogger())
	require.NotNil(t, err)
	require.NotNil(t, base)

	lolClient, err := NewClient(base)
	require.NotNil(t, err)
	require.NotNil(t, lolClient)

	data, err := lolClient.Summoner.GatherDataByName("this is a really long name that can't exist")
	require.NotNil(t, err)
	require.Nil(t, data)
}

func TestSummonerClient_AnalyzeData(t *testing.T) {

}

func TestSummonerClient_AnalyzeData_Nil_Error(t *testing.T) {
	config, err := configuration.GetConfig(configFile)
	require.Nil(t, err)
	require.NotNil(t, config)

	base, err := internal.NewClient(config.APIKey, api.Region(config.Region), logrus.StandardLogger())
	require.NotNil(t, err)
	require.NotNil(t, base)

	lolClient, err := NewClient(base)
	require.NotNil(t, err)
	require.NotNil(t, lolClient)

	analysis, err := lolClient.Summoner.AnalyzeData(nil)
	require.NotNil(t, err)
	require.Nil(t, analysis)
}

func TestSummonerClient_GenerateSummonerReport(t *testing.T) {

}
