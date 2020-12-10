package goScout

import (
	"github.com/matcarter/goScout/internal/configuration"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewScout(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "Good case: config file found",
			path:    "internal/configuration/init.json",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := configuration.GetConfig(tt.path)
			require.Nil(t, err)

			scout, err := NewScout(config.APIKey, config.Region)
			if tt.wantErr {
				require.NotNil(t, err)
				require.Nil(t, scout)
			} else {
				require.Nil(t, err)
				require.NotNil(t, scout)
			}
		})
	}
}
