package lol

import "github.com/matcarter/goScout/internal"

// MatchClient allows access to the internal client
type MatchClient struct {
	c *internal.Client
}
