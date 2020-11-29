package lol

import (
	"github.com/matcarter/goScout/api"
	"github.com/matcarter/goScout/internal"
	"github.com/sirupsen/logrus"
)

// SummonerClient allows access to the internal client
type SummonerClient struct {
	c *internal.Client
}

// ScoutByName takes a summoner name as a string and creates a SummonerReport from the raw data pulled from the Riot
// API tied to that summoner name
func (s *SummonerClient) ScoutByName(name string) (*SummonerReport, error) {
	// Gather raw data from the Riot API
	summonerData, err := s.gatherDataByName(name)
	if err != nil {
		s.logger().WithFields(logrus.Fields{
			"method": "ScoutByName",
			"name":   name,
			"err":    err,
		}).Warn("Failed to gather data for summoner")

		return nil, err
	}
	s.logger().WithFields(logrus.Fields{
		"name": name,
	}).Info("Successfully gathered data for summoner")

	// Analyze the raw data into meaningful data
	analyzedData, err := s.AnalyzeData(summonerData)
	if err != nil {
		s.logger().WithFields(logrus.Fields{
			"method": "ScoutByName",
			"name":   name,
			"err":    err,
		}).Warn("Failed to analyze summoner data")

		return nil, err
	}
	s.logger().WithFields(logrus.Fields{
		"name": name,
	}).Info("Successfully analyzed summoner data")

	// Format the analyzed data in a meaningful way for the user
	summonerReport, err := s.GenerateSummonerReport(analyzedData)
	if err != nil {
		s.logger().WithFields(logrus.Fields{
			"method": "ScoutByName",
			"name":   name,
			"err":    err,
		}).Warn("Failed to generate report for summoner")

		return nil, err
	}
	s.logger().WithFields(logrus.Fields{
		"name": name,
	}).Info("Successfully generated report for summoner")

	return summonerReport, nil
}

// gatherDataByName pulls all relevant data from the Riot API given a name and stores it within a SummonerData
// struct that is then returned
func (s *SummonerClient) gatherDataByName(name string) (*SummonerData, error) {
	// Query the Riot API's GetByName method to get account information relating to said name
	summoner, err := s.c.Golio.Riot.LoL.Summoner.GetByName(name)
	if err != nil {
		s.logger().WithFields(logrus.Fields{
			"method": "gatherDataByName",
			"name":   name,
			"err":    err,
		}).Warn("Failed to fetch summoner")
		return nil, err
	}
	s.logger().WithFields(logrus.Fields{
		"name": name,
	}).Debug("Successfully fetched summoner")

	summonerData := &SummonerData{
		Summoner: summoner,
	}

	// Query the Riot API's ChampionMastery List method to get all champions and their mastery level
	masteries, err := s.c.Golio.Riot.LoL.ChampionMastery.List(summoner.ID)
	if err != nil {
		s.logger().WithFields(logrus.Fields{
			"method": "gatherDataByName",
			"name":   name,
			"err":    err,
		}).Warn("Failed to fetch summoner masteries")
		return nil, err
	}
	s.logger().WithFields(logrus.Fields{
		"name": name,
	}).Debug("Successfully fetched masteries for summoner")

	summonerData.Mastery = masteries

	// Query the Riot API's Match List method to get the last 100 games they played
	matches, err := s.c.Golio.Riot.LoL.Match.List(summoner.AccountID, 0, 100)
	if err != nil {
		s.logger().WithFields(logrus.Fields{
			"method": "gatherDataByName",
			"name":   name,
			"err":    err,
		}).Warn("Failed to fetch summoner matches")
		return nil, err
	}
	s.logger().WithFields(logrus.Fields{
		"name": name,
	}).Debug("Successfully fetched matches for summoner")

	summonerData.Matches = matches

	return summonerData, nil
}

// AnalyzeData analyzes the raw data stored within a SummonerData struct, finds key information within the raw data,
// and returns it in the form of a SummonerDataAnalysis struct
func (s *SummonerClient) AnalyzeData(summonerData *SummonerData) (*SummonerDataAnalysis, error) {
	if summonerData == nil {
		s.logger().WithFields(logrus.Fields{
			"method": "AnalyzeData",
		}).Warn("Cannot analyze nil data")

		return nil, api.ErrNilSummonerData
	}

	analysis := &SummonerDataAnalysis{
		Summoner:       summonerData.Summoner,
		HighestMastery: summonerData.Mastery[:10],
	}

	return analysis, nil
}

// GenerateSummonerReport formats the given SummonerDataAnalysis in a way meaningful to the user
func (s *SummonerClient) GenerateSummonerReport(analyzedData *SummonerDataAnalysis) (*SummonerReport, error) {
	return nil, nil
}

// logger prepends the logger for the summoner module with a named tag
func (s *SummonerClient) logger() logrus.FieldLogger {
	return s.c.Logger().WithField("category", "summoner")
}
