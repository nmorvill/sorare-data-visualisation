package sorare

import (
	"context"
	"fmt"
	"os"

	"github.com/machinebox/graphql"
)

func getDomesticLeagueSlugs() []string {
	competitions, err := callSorareApi[OpenLeagues](graphql.NewRequest(openLeaguesQuery))
	if err != nil {
		panic(err)
	}
	var ret []string
	for _, c := range competitions.Football.LeaguesOpenForGameStats {
		if c.Format == "DOMESTIC_LEAGUE" {
			ret = append(ret, c.Slug)
		}
	}
	return ret
}

func getClubsSlugs(slug string) []string {
	league, err := callSorareApi[League](graphql.NewRequest(fmt.Sprintf(leagueQuery, slug)))
	if err != nil {
		panic(err)
	}
	var ret []string
	for _, c := range league.Football.Competition.Clubs.Nodes {
		ret = append(ret, c.Slug)
	}
	return ret
}

func getPlayersSlugs(slug string) []PlayerName {
	team, err := callSorareApi[Team](graphql.NewRequest(fmt.Sprintf(teamQuery, slug)))
	if err != nil {
		panic(err)
	}
	var ret []PlayerName
	for _, p := range team.Football.Club.ActivePlayers.Nodes {
		ret = append(ret, PlayerName{DisplayName: p.DisplayName, Slug: p.Slug})
	}
	return ret
}

func GetPlayerInformations(slug string) (Player, error) {
	return callSorareApi[Player](graphql.NewRequest(fmt.Sprintf(playersInfoQuery, slug)))
}

func callSorareApi[K interface{}](req *graphql.Request) (K, error) {
	api_key := os.Getenv("SORARE_API_KEY")
	client := graphql.NewClient("https://api.sorare.com/graphql")
	req.Header.Set("APIKEY", api_key)
	var ret K
	if err := client.Run(context.Background(), req, &ret); err != nil {
		return ret, err
	}
	return ret, nil
}
