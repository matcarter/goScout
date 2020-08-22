package riot

import (
	"github.com/KnutZuidema/golio/api"
	"github.com/matcarter/goScout/internal"
	"github.com/matcarter/goScout/riot/lol"
	"github.com/sirupsen/logrus"
)

type Client struct {
	LoL *lol.Client
}

func NewClient(apiKey string, region api.Region, logger logrus.FieldLogger) *Client {
	baseClient := internal.NewClient(apiKey, region, logger)
	c := &Client{
		LoL: lol.NewClient(baseClient),
	}

	return c
}
