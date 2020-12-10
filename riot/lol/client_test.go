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
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "Good case: config file found",
			path:    "../../internal/configuration/init.json",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := configuration.GetConfig(tt.path)
			require.Nil(t, err)
			require.NotNil(t, config)

			base, err := internal.NewClient(config.APIKey, api.Region(config.Region), logrus.StandardLogger())
			require.Nil(t, err)
			require.NotNil(t, base)

			client, err := NewClient(base)
			if tt.wantErr {
				require.NotNil(t, err)
				require.Nil(t, client)
			} else {
				require.Nil(t, err)
				require.NotNil(t, client)
			}
		})
	}
}
