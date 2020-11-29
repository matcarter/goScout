package lol

import "github.com/matcarter/goScout/internal"

// Client is the struct for the client containing all available interfaces for interacting with the
// League of Legends API provided by Riot Games
type Client struct {
	Summoner *SummonerClient
}

// NewClient initializes a new League of Legends client given a setup internal client as it's base
func NewClient(base *internal.Client) (c *Client, err error) {
	c = &Client{
		Summoner: &SummonerClient{c: base},
	}

	return c, err
}
