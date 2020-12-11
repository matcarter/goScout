package lol

import "github.com/matcarter/goScout/internal"

// SummonerClient allows access to the internal client
type SummonerClient struct {
	c *internal.Client
}
