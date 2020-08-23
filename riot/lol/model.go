package lol

import "github.com/KnutZuidema/golio/riot/lol"

type ScoutOptions struct {
	NumberOfMasteries int
	NumberOfGames     int
	ExcludedQueues    []int
}

type SummonerData struct {
	Summoner *lol.Summoner
	Mastery  []*lol.ChampionMastery
	Matches  *lol.Matchlist
}

type SummonerDataAnalysis struct {
	Summoner *lol.Summoner
}

type SummonerReport struct {
}
