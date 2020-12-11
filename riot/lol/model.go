package lol

import (
	"fmt"

	"github.com/KnutZuidema/golio/riot/lol"
)

// ScoutOptions defines basic optional values that can be configured when scouting users
type ScoutOptions struct {
	NumberOfMastery int   // NumberOfMastery specifies the upper limit of number of champion mastery values to include
	NumberOfGames   int   // NumberOfGames specifies the upper limit of how many games to be analyze
	ExcludedQueues  []int // ExcludedQueues specifies what queues should not be included in the report
}

// SummonerData defines the high-level information that can be pulled from the Riot API about a player (summoner)
type SummonerData struct {
	Summoner *lol.Summoner          // Summoner contains account information for the summoner
	Mastery  []*lol.ChampionMastery // Mastery is a list of data, champions, and the summoner's mastery on said champion
	Matches  *lol.Matchlist         // Matches is a collection of games play and granular data for each game
}

// GetHighestMastery returns the data about which champion the summoner has the highest mastery for
func (sd *SummonerData) GetHighestMastery(n int) (highestMastery []*lol.ChampionMastery, err error) {
	if n < 0 {
		return nil, fmt.Errorf("error execpted n [%d] to be greater than 0", n)
	}

	if n > len(sd.Mastery) {
		return nil, fmt.Errorf("error expected n [%d] to be less than or equal to mastery [%d]", n, len(sd.Mastery))
	}

	return sd.Mastery[:n], nil
}

// SummonerDataAnalysis outlines hand-crafted data points collected through the analysis of the raw data contained
// within SummonerData
type SummonerDataAnalysis struct {
	Summoner                  *lol.Summoner          // Summoner contains account information for the summoner
	HighestRank               string                 // HighestRank is the highest rank achieved within ranked/competitive modes
	CurrentRank               string                 // CurrentRank is the current rank within ranked/competitive modes
	HighestMastery            []*lol.ChampionMastery // HighestMastery is a list of champions and their mastery level
	RecentlyPlayedChampions   map[int]int            // RecentlyPlayedChampions is a map of champion IDs to number of games played recently
	FrequentlyPlayedChampions map[int]int            // FrequentlyPlayedChampions is a map of champion IDs to number of games played within a timeframe
	RecentlyPlayedLanes       map[string]int         // RecentlyPlayedLanes is a map of lane names to number of games played in that lane recently
	FrequentlyPlayedLanes     map[string]int         // FrequentlyPlayedLanes is a map of lane names to number of games played in that lane within a timeframe
}

// SummonerReport outlines the final report of the analyzed data within SummonerDataAnalysis
type SummonerReport struct {
	Name        string
	Level       int
	LastUpdated int
}
