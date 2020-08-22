package main

import (
	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
	"github.com/sirupsen/logrus"
)

type Scout struct {
	apiKey string
	client *golio.Client
}

func NewScout(apiKey string) *Scout {
	s := &Scout{
		apiKey: apiKey,
		client: golio.NewClient(apiKey,
			golio.WithRegion(api.RegionNorthAmerica),
			golio.WithLogger(logrus.New().WithField("foo", "bar"))),
	}

	return s
}
