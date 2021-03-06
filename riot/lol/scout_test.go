package lol

import (
	"testing"

	"github.com/KnutZuidema/golio/api"
	"github.com/matcarter/goScout/internal"
	"github.com/matcarter/goScout/internal/configuration"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

var configFile = "../../internal/configuration/init.json"

func TestScoutClient_ScoutByName(t *testing.T) {
	tests := []struct {
		name     string
		summoner string
		wantErr  bool
	}{
		{
			name:    "Bad case: nil summoner name",
			wantErr: true,
		},
		{
			name:     "Bad case: no summoner name provided",
			summoner: "",
			wantErr:  true,
		},
		{
			name:     "Good case: summoner exists in config region",
			summoner: "Bully",
			wantErr:  false,
		},
		{
			name:     "Bad case: summoner name is too long",
			summoner: "this summoner name is way too long and does not exist",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := configuration.GetConfig(configFile)
			require.Nil(t, err)
			require.NotNil(t, config)

			base, err := internal.NewClient(config.APIKey, api.Region(config.Region), logrus.StandardLogger())
			require.Nil(t, err)
			require.NotNil(t, base)

			lolClient, err := NewClient(base)
			require.Nil(t, err)
			require.NotNil(t, lolClient)

			report, err := lolClient.Scout.ScoutByName(tt.summoner)
			if tt.wantErr {
				require.NotNil(t, err)
				require.Nil(t, report)
			} else {
				require.Nil(t, err)
				require.NotNil(t, report)
			}
		})
	}
}

func TestScoutClient_GatherDataByName(t *testing.T) {
	tests := []struct {
		name     string
		summoner string
		wantErr  bool
	}{
		{
			name:    "Bad case: nil summoner name",
			wantErr: true,
		},
		{
			name:     "Bad case: no summoner name provided",
			summoner: "",
			wantErr:  true,
		},
		{
			name:     "Good case: summoner exists in config region",
			summoner: "Bully",
			wantErr:  false,
		},
		{
			name:     "Bad case: summoner name is too long",
			summoner: "this summoner name is way too long and does not exist",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := configuration.GetConfig(configFile)
			require.Nil(t, err)
			require.NotNil(t, config)

			base, err := internal.NewClient(config.APIKey, api.Region(config.Region), logrus.StandardLogger())
			require.Nil(t, err)
			require.NotNil(t, base)

			lolClient, err := NewClient(base)
			require.Nil(t, err)
			require.NotNil(t, lolClient)

			data, err := lolClient.Scout.GatherDataByName(tt.summoner)
			if tt.wantErr {
				require.NotNil(t, err)
				require.Nil(t, data)
			} else {
				require.Nil(t, err)
				require.NotNil(t, data)
			}
		})
	}
}

func TestScoutClient_AnalyzeData(t *testing.T) {
	tests := []struct {
		name         string
		summonerData *SummonerData
		wantErr      bool
	}{
		{
			name:         "Bad case: nil summoner data",
			summonerData: nil,
			wantErr:      true,
		},
		{
			name:         "Good case: non-nil summoner data",
			summonerData: &SummonerData{},
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := configuration.GetConfig(configFile)
			require.Nil(t, err)
			require.NotNil(t, config)

			base, err := internal.NewClient(config.APIKey, api.Region(config.Region), logrus.StandardLogger())
			require.Nil(t, err)
			require.NotNil(t, base)

			lolClient, err := NewClient(base)
			require.Nil(t, err)
			require.NotNil(t, lolClient)

			analysis, err := lolClient.Scout.AnalyzeData(tt.summonerData)
			if tt.wantErr {
				require.NotNil(t, err)
				require.Nil(t, analysis)
			} else {
				require.Nil(t, err)
				require.NotNil(t, analysis)
			}
		})
	}
}

func TestScoutClient_GenerateSummonerReport(t *testing.T) {
	tests := []struct {
		name         string
		analyzedData *ScoutAnalysis
		wantErr      bool
	}{
		{
			name:         "Bad case: nil summoner data",
			analyzedData: nil,
			wantErr:      true,
		},
		{
			name:         "Good case: non-nil summoner data",
			analyzedData: &ScoutAnalysis{},
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := configuration.GetConfig(configFile)
			require.Nil(t, err)
			require.NotNil(t, config)

			base, err := internal.NewClient(config.APIKey, api.Region(config.Region), logrus.StandardLogger())
			require.Nil(t, err)
			require.NotNil(t, base)

			lolClient, err := NewClient(base)
			require.Nil(t, err)
			require.NotNil(t, lolClient)

			analysis, err := lolClient.Scout.GenerateSummonerReport(tt.analyzedData)
			if tt.wantErr {
				require.NotNil(t, err)
				require.Nil(t, analysis)
			} else {
				require.Nil(t, err)
				require.NotNil(t, analysis)
			}
		})
	}
}
