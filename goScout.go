package goScout

import (
	"github.com/KnutZuidema/golio/api"
	"github.com/matcarter/goScout/riot"
	"github.com/sirupsen/logrus"
)

// Scout outlines necessary data for goScout to properly function
type Scout struct {
	apiKey string             // apiKey is the Riot Games API key used to pull data from their services
	logger logrus.FieldLogger // logger is logging object used within the Scout object
	region api.Region         // region is the server region from which the Scout pulls data from
	riot   *riot.Client       // riot is a pointer to the client for the Riot games service
}

// NewScout initializes a new
func NewScout(apiKey string, region string) (s *Scout, err error) {
	s = &Scout{
		apiKey: apiKey,
		logger: logrus.StandardLogger(),
		region: api.Region(region),
	}

	s.riot, err = riot.NewClient(s.apiKey, s.region, s.logger)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"method": "NewScout",
			"apiKey": apiKey,
			"region": region,
			"err":    err,
		}).Warn("Failed to initialize the Scout struct")
	}

	return s, nil
}
