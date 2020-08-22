package main

import (
	"fmt"
	_ "net/http"

	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/riot/lol"
	"github.com/sirupsen/logrus"
)

func main() {
	client := golio.NewClient("RGAPI-90d213a3-0ca1-41bc-b091-66d40e525cf1",
		golio.WithRegion(api.RegionNorthAmerica),
		golio.WithLogger(logrus.New().WithField("foo", "bar")))
	summoner, _ := client.Riot.LoL.Summoner.GetByName("Bully")

	fmt.Printf("%s is a level %d summoner\n", summoner.Name, summoner.SummonerLevel)

	champion, _ := client.DataDragon.GetChampion("Katarina")
	mastery, err := client.Riot.LoL.ChampionMastery.Get(summoner.ID, champion.Key)

	if err != nil {
		fmt.Printf("%s has not played any games on %s\n", summoner.Name, champion.Name)
	} else {
		fmt.Printf("%s has mastery level %d with %d points on %s\n", summoner.Name, mastery.ChampionLevel,
			mastery.ChampionPoints, champion.Name)
	}

	challengers, _ := client.Riot.LoL.League.GetChallenger(lol.QueueRankedSolo)
	rank1 := challengers.GetRank(0)

	fmt.Printf("%s is the highest ranked player with %d league points\n", rank1.SummonerName, rank1.LeaguePoints)
}
