package lol

import (
	"github.com/KnutZuidema/golio/riot/lol"
)

type ScoutOptions struct {
	NumberOfMastery int
	NumberOfGames   int
	ExcludedQueues  []int
}

type SummonerData struct {
	Summoner *lol.Summoner
	Mastery  []*lol.ChampionMastery
	Matches  *lol.Matchlist
}

func (sd *SummonerData) GetHighestMastery(n int) []*lol.ChampionMastery {
	return sd.Mastery[:n]
}

type SummonerDataAnalysis struct {
	Summoner                  *lol.Summoner
	HighestRank               string
	CurrentRank               string
	HighestMastery            []*lol.ChampionMastery
	RecentlyPlayedChampions   map[int]int
	FrequentlyPlayedChampions map[int]int
	RecentlyPlayedLanes       map[string]int
	FrequentlyPlayedLanes     map[string]int
}

type SummonerReport struct {
}
