package internal

import (
	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
	"github.com/sirupsen/logrus"
)

type Client struct {
	APIKey string
	Golio  *golio.Client
	L      logrus.FieldLogger
	Region api.Region
}

func NewClient(apiKey string, region api.Region, logger logrus.FieldLogger) *Client {
	c := &Client{
		APIKey: apiKey,
		L:      logger,
		Region: region,
	}

	c.Golio = golio.NewClient(apiKey,
		golio.WithRegion(c.Region),
		golio.WithLogger(logrus.New().WithField("method", "NewClient")))

	return c
}

func (c *Client) Logger() logrus.FieldLogger {
	return c.L.WithField("region", c.Region)
}
