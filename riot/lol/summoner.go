package lol

import (
	"github.com/matcarter/goScout/internal"
	"github.com/sirupsen/logrus"
)

type SummonerClient struct {
	c *internal.Client
}

func (s *SummonerClient) ScoutByName(name string) (*SummonerReport, error) {
	summoner, _ := s.c.Golio.Riot.LoL.Summoner.GetByName(name)

	report := &SummonerReport{
		Name:  summoner.Name,
		Level: summoner.SummonerLevel,
	}

	return report, nil
}

func (s *SummonerClient) logger() logrus.FieldLogger {
	return s.c.Logger().WithField("category", "summoner")
}
