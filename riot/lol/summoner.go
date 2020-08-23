package lol

import (
	"github.com/matcarter/goScout/internal"
	"github.com/sirupsen/logrus"
)

type SummonerClient struct {
	c *internal.Client
}

func (s *SummonerClient) ScoutByName(name string) (*SummonerReport, error) {
	summonerData, err := s.gatherDataByName(name)
	if err != nil {
		s.logger().WithFields(logrus.Fields{
			"name": name,
			"err":  err,
		}).Warn("Failed to gather data for summoner")
		return nil, err
	}
	s.logger().WithFields(logrus.Fields{
		"name": name,
	}).Info("Successfully gathered data for summoner")

	analyzedData, err := s.AnalyzeData(summonerData)
	if err != nil {
		s.logger().WithFields(logrus.Fields{
			"name": name,
			"err":  err,
		}).Warn("Failed to analyze summoner data")
		return nil, err
	}
	s.logger().WithFields(logrus.Fields{
		"name": name,
	}).Info("Successfully analyzed summoner data")

	summonerReport, err := s.GenerateSummonerReport(analyzedData)
	if err != nil {
		s.logger().WithFields(logrus.Fields{
			"name": name,
			"err":  err,
		}).Warn("Failed to generate report for summoner")
		return nil, err
	}
	s.logger().WithFields(logrus.Fields{
		"name": name,
	}).Info("Successfully generated report for summoner")

	return summonerReport, nil
}

func (s *SummonerClient) gatherDataByName(name string) (*SummonerData, error) {
	summoner, err := s.c.Golio.Riot.LoL.Summoner.GetByName(name)
	if err != nil {
		s.logger().WithFields(logrus.Fields{
			"name": name,
			"err":  err,
		}).Warn("Failed to fetch summoner")
		return nil, err
	}
	s.logger().WithFields(logrus.Fields{
		"name": name,
	}).Debug("Successfully fetched summoner")

	summonerData := &SummonerData{
		Summoner: summoner,
	}

	masteries, err := s.c.Golio.Riot.LoL.ChampionMastery.List(summoner.ID)
	if err != nil {
		s.logger().WithFields(logrus.Fields{
			"name": name,
			"err":  err,
		}).Warn("Failed to fetch summoner masteries")
		return nil, err
	}
	s.logger().WithFields(logrus.Fields{
		"name": name,
	}).Debug("Successfully fetched masteries for summoner")

	summonerData.Mastery = masteries

	matches, err := s.c.Golio.Riot.LoL.Match.List(summoner.AccountID, 0, 100)
	if err != nil {
		s.logger().WithFields(logrus.Fields{
			"name": name,
			"err":  err,
		}).Warn("Failed to fetch summoner matches")
		return nil, err
	}
	s.logger().WithFields(logrus.Fields{
		"name": name,
	}).Debug("Successfully fetched matches for summoner")

	summonerData.Matches = matches

	return summonerData, nil
}

func (s *SummonerClient) AnalyzeData(summonerData *SummonerData) (*SummonerDataAnalysis, error) {
	return nil, nil
}

func (s *SummonerClient) GenerateSummonerReport(analyzedData *SummonerDataAnalysis) (*SummonerReport, error) {
	return nil, nil
}

func (s *SummonerClient) logger() logrus.FieldLogger {
	return s.c.Logger().WithField("category", "summoner")
}
