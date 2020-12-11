package lol

import "github.com/matcarter/goScout/internal"

// Client is the struct for the client containing all available interfaces for interacting with the
// League of Legends API provided by Riot Games
type Client struct {
	Champion *ChampionClient // Champion is the client for gathering information about champions
	Match    *MatchClient    // Match is the client for gathering information about matches
	Scout    *ScoutClient    // Scout is the client for analyzing player data and crafting a report
	Summoner *SummonerClient //Summoner is the client for gathering information about summoners
}

// NewClient initializes a new League of Legends client given a setup internal client as it's base
func NewClient(base *internal.Client) (c *Client, err error) {
	c = &Client{
		Champion: &ChampionClient{c: base},
		Match:    &MatchClient{c: base},
		Scout:    &ScoutClient{c: base},
		Summoner: &SummonerClient{c: base},
	}

	return c, err
}
