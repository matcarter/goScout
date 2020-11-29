package riot

import (
	"github.com/KnutZuidema/golio/api"
	"github.com/matcarter/goScout/internal"
	"github.com/matcarter/goScout/riot/lol"
	"github.com/sirupsen/logrus"
)

// Client provides access to a sub-clients that interact with various Riot Games APIs for their available games
type Client struct {
	LoL *lol.Client // LoL is the name of the League of Legends client found in the "lol" subpackage
}

// NewClient initializes the client for use
func NewClient(apiKey string, region api.Region, logger logrus.FieldLogger) (c *Client, err error) {
	// Initialize a new internal client given the API Key, Region, and a logger
	baseClient, err := internal.NewClient(apiKey, region, logger)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"method": "NewClient",
			"err":    err,
		}).Warn("Failed to initialize an internal client")
	}

	// Initialize a new League of Legends external client given the internal client
	lolClient, err := lol.NewClient(baseClient)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"method": "NewClient",
			"err":    err,
		}).Warn("Failed to initialize an external League of Legends client")
	}

	// Initialize a new external client
	c = &Client{
		LoL: lolClient,
	}

	return c, nil
}
