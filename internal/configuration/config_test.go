package configuration

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetConfig(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "Bad case: config file not found",
			path:    "definitely_not_a_config_file.json",
			wantErr: true,
		},
		{
			name:    "Good case: config file found",
			path:    "init.json",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := GetConfig(tt.path)
			if tt.wantErr {
				require.NotNil(t, err)
				require.Nil(t, config)
			} else {
				require.Nil(t, err)
				require.NotNil(t, config)
			}
		})
	}
}
