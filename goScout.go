package goScout

import (
	"github.com/KnutZuidema/golio/api"
	"github.com/matcarter/goScout/riot"
	"github.com/sirupsen/logrus"
)

type Scout struct {
	apiKey string
	logger logrus.FieldLogger
	region api.Region
	riot   *riot.Client
}

func NewScout(apiKey string, region string) *Scout {
	s := &Scout{
		apiKey: apiKey,
		logger: logrus.StandardLogger(),
		region: api.Region(region),
	}

	s.riot = riot.NewClient(s.apiKey, s.region, s.logger)

	return s
}
