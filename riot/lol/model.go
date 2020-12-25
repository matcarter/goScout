package lol

import (
	"fmt"

	"github.com/KnutZuidema/golio/riot/lol"
)

/*
 * Summoner
 */

// SummonerData defines the high-level information that can be pulled from the Riot API about a player (summoner)
type SummonerData struct {
	Summoner *lol.Summoner          // Summoner contains account information for the summoner
	Mastery  []*lol.ChampionMastery // Mastery is a list of data, champions, and the summoner's mastery on said champion
	Matches  *lol.Matchlist         // Matches is a collection of games play and granular data for each game
	Leagues  []*lol.LeagueItem      // Leagues is a list of leagues for the summoner
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

// TODO: add fields relevant to just the summoner
// account age, is the account a smurf, more?
type SummonerAnalysis struct {
	Summoner *lol.Summoner
}

/*
 * Champion
 */

// TODO: add fields relevant to a summoner's champions
// highest mastery, champ preference (mages, assassins, etc.), K/D, CS, WR
type ChampionAnalysis struct {
}

/*
 * Match
 */

// TODO: add fields relevant to a summoner's matches
// players recently queued with, most recent W/L ratio, recent lanes, recent champions, recent K/D, etc.
type MatchAnalysis struct {
}

/*
 * League
 */

// TODO: add fields relevant to a summoner's leagues
// just lol.LeagueItem info?
type LeagueAnalysis struct {
}

/*
 * Scout
 */

// ScoutOptions defines basic optional values that can be configured when scouting users
type ScoutOptions struct {
	NumberOfMastery int   // NumberOfMastery specifies the upper limit of number of champion mastery values to include
	NumberOfGames   int   // NumberOfGames specifies the upper limit of how many games to be analyze
	ExcludedQueues  []int // ExcludedQueues specifies what queues should not be included in the report
}

// ScoutAnalysis outlines hand-crafted data points collected through the analysis of the raw data contained
// within SummonerData
type ScoutAnalysis struct {
	SummonerAnalysis
	ChampionAnalysis
	MatchAnalysis
	LeagueAnalysis
}

// ScoutReport outlines the final report of the analyzed data within ScoutAnalysis
type ScoutReport struct {
}
