package internal

import (
	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
	"github.com/sirupsen/logrus"
)

// Client is the internal facing client that possess all access to the golio library
type Client struct {
	APIKey string             // APIKey is the API Key used to authenticate with the Riot Games API
	Golio  *golio.Client      // Golio is a pointer to the client setup from the golio library
	L      logrus.FieldLogger // L is a logger used within the client
	Region api.Region         // Region is a golio API region specifying what region the golio client is configured for
}

// NewClient initializes a new internal client and configures golio with the specified information
func NewClient(apiKey string, region api.Region, logger logrus.FieldLogger) (c *Client, err error) {
	c = &Client{
		APIKey: apiKey,
		L:      logger,
		Region: region,
	}

	c.Golio = golio.NewClient(apiKey,
		golio.WithRegion(c.Region),
		golio.WithLogger(logrus.New().WithField("method", "NewClient")))

	return c, nil
}

// Logger is the method used to prepend a logger with the region being used for API access
func (c *Client) Logger() (logger logrus.FieldLogger) {
	return c.L.WithField("region", c.Region)
}
