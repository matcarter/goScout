package lol

import "github.com/matcarter/goScout/internal"

type Client struct {
	Summoner *SummonerClient
}

func NewClient(base *internal.Client) *Client {
	return &Client{
		Summoner: &SummonerClient{c: base},
	}
}
