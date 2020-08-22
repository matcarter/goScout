package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewScout(t *testing.T) {
	scout := NewScout("api_key")

	require.NotNil(t, scout)
}
