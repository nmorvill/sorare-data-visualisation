package players

import (
	"bytes"
	"fmt"
	"html/template"
	"sg/internal/sorare"
	"sg/internal/utils"
	"strings"
)

type playerPage struct {
	Infos     playerInfos
	MainStats playerMainStats
	Tags      []playerTag
	Games     []playerGame
}

type playerInfos struct {
	Picture        string
	Name           string
	Position       string
	Age            int
	CountryName    string
	CountryPicture string
	ClubName       string
	ClubPicture    string
}

type playerMainStats struct {
	L5Avg       int
	L5Presence  int
	L5Color     string
	L15Avg      int
	L15Presence int
	L15Color    string
	L50Avg      int
	L50Presence int
	L50Color    string
}

type playerTag struct {
	Title string
	Color string
}

type playerGame struct {
	Played          bool
	Score           float32
	Color           string
	HomeTeamPicture string
	HomeTeamScore   int
	AwayTeamPicture string
	AwayTeamScore   int
	Gameweek        int
	Categories      []playerCategory
}

type playerCategory struct {
	Score float32
	Name  string
	Color string
}

type tagsColors string

const (
	DARK_RED   tagsColors = "#CD4115"
	RED        tagsColors = "#D75E5E"
	YELLOW     tagsColors = "#F1AF4B"
	GREEN      tagsColors = "#5EA258"
	DARK_GREEN tagsColors = "#34732F"
)

func RenderPlayerPlage(slug string) []byte {
	t, err := template.ParseFiles("./internal/players/playerPage.tmpl")
	if err != nil {
		fmt.Println(err.Error())
	}
	var out bytes.Buffer
	err = t.Execute(&out, slug)
	if err != nil {
		fmt.Println(err.Error())
	}
	return out.Bytes()
}

func RenderGeneralPlayerPage(slug string) []byte {
	var out bytes.Buffer
	player, err := sorare.GetPlayerInformations(slug)
	if err != nil {
		out.WriteString("<h2>Player not found !</h2>")
		return out.Bytes()
	}
	page := getPlayerPage(player)
	t, err := template.ParseFiles("./internal/players/generalPlayerPage.tmpl")
	if err != nil {
		panic(err)
	}
	err = t.Execute(&out, page)
	if err != nil {
		panic(err)
	}
	return out.Bytes()
}

func RenderSearchResults(search string, names *([]sorare.PlayerName)) []byte {
	search = strings.ToLower(search)
	var out bytes.Buffer
	var matches []sorare.PlayerName
	for _, p := range *names {
		if strings.Contains(strings.ToLower(p.DisplayName), search) {
			matches = append(matches, p)
		}
	}
	t, err := template.ParseFiles("./internal/players/searchResult.tmpl")
	if err != nil {
		panic(err)
	}
	if len(matches) > 5 {
		matches = matches[:5]
	}
	err = t.Execute(&out, matches)
	if err != nil {
		panic(err)
	}
	return out.Bytes()
}

func getPlayerPage(p sorare.Player) playerPage {
	var page playerPage
	page.Infos = getInfos(p)
	page.Tags = getTags(p.Football.Player.Slug)
	page.MainStats, page.Games = getPlayerStats(p)
	return page
}

func getInfos(p sorare.Player) playerInfos {
	var header playerInfos
	header.Name = p.Football.Player.DisplayName
	header.Position = p.Football.Player.Position
	header.Age = p.Football.Player.Age
	header.CountryName = p.Football.Player.Country.ThreeLetterCode
	header.CountryPicture = p.Football.Player.Country.FlagUrl
	header.ClubPicture = p.Football.Player.ActiveClub.PictureUrl
	header.ClubName = p.Football.Player.ActiveClub.ShortName
	header.Picture = p.Football.Player.PictureUrl
	return header
}

func getPlayerStats(p sorare.Player) (playerMainStats, []playerGame) {
	var l5, l15, l50 float32
	var l5M, l15M, l50M int
	var l5P, l15P, l50P float32
	var games []playerGame
	for i, s := range p.Football.Player.AllSo5Scores.Nodes {
		if s.PlayerGameStats.MinsPlayed > 0 {
			if i < 5 {
				l5 += s.Score
				l5M += s.PlayerGameStats.MinsPlayed
				l5P++
			}
			if i < 15 {
				l15 += s.Score
				l15M += s.PlayerGameStats.MinsPlayed
				l15P++
			}
			l50 += s.Score
			l50M += s.PlayerGameStats.MinsPlayed
			l50P++
		}
		games = append(games, getGameStats(s, p.Football.Player.Position))
	}
	l5 /= l5P
	l15 /= l15P
	l50 /= l50P
	l5M = int(float32(l5M) * 100 / float32(5*90))
	l15M = int(float32(l15M) * 100 / float32(15*90))
	l50M = int(float32(l50M) * 100 / float32(50*90))
	stats := playerMainStats{
		L5Avg:       int(l5),
		L5Presence:  l5M,
		L5Color:     utils.GetColorCodeOfNote(int(l5)),
		L15Avg:      int(l15),
		L15Presence: l15M,
		L15Color:    utils.GetColorCodeOfNote(int(l15)),
		L50Avg:      int(l50),
		L50Presence: l50M,
		L50Color:    utils.GetColorCodeOfNote(int(l50)),
	}
	return stats, games
}

func getGameStats(s sorare.Score, pos string) playerGame {
	var g playerGame
	g.HomeTeamPicture = s.Game.HomeTeam.PictureUrl
	g.HomeTeamScore = s.Game.HomeGoals
	g.AwayTeamPicture = s.Game.AwayTeam.PictureUrl
	g.AwayTeamScore = s.Game.AwayGoals
	g.Gameweek = utils.GetGameweekFromString(s.Game.Date)
	if s.PlayerGameStats.MinsPlayed == 0 {
		g.Played = false
		return g
	}
	g.Score = s.Score
	g.Color = utils.GetColorCodeOfNote(int(s.Score))
	g.Played = true
	goalkeeping := playerCategory{Name: "Goalkeeping"}
	defending := playerCategory{Name: "Defending"}
	passing := playerCategory{Name: "Passing"}
	possession := playerCategory{Name: "Possession"}
	attacking := playerCategory{Name: "Attacking"}
	for _, st := range s.DetailedScore {
		switch st.Category {
		case "GOALKEEPING":
			goalkeeping.Score += st.TotalScore
			break
		case "DEFENDING":
			defending.Score += st.TotalScore
			break
		case "POSSESSION":
			possession.Score += st.TotalScore
			break
		case "PASSING":
			passing.Score += st.TotalScore
			break
		case "ATTACKING":
			attacking.Score += st.TotalScore
			break
		case "GENERAL":
			if st.Stat == "was_fouled" {
				possession.Score += st.TotalScore
			} else {
				defending.Score += st.TotalScore
			}
			break
		}
	}
	var categories []playerCategory
	if pos != "Goalkeeper" {
		categories = []playerCategory{defending, possession, passing, attacking}
	} else {
		categories = []playerCategory{goalkeeping, possession, passing}
	}
	for i, c := range categories {
		if c.Score > 10 {
			categories[i].Color = string(DARK_GREEN)
		} else if c.Score > 2 {
			categories[i].Color = string(GREEN)
		} else if c.Score > -2 {
			categories[i].Color = string(YELLOW)
		} else if c.Score > -10 {
			categories[i].Color = string(RED)
		} else {
			categories[i].Color = string(DARK_RED)
		}
	}
	g.Categories = categories
	return g
}

/* TODO */
func getTags(slug string) []playerTag {
	return []playerTag{
		{Title: "Bad at home", Color: string(RED)},
		{Title: "Often misses", Color: string(RED)},
		{Title: "Low variance", Color: string(YELLOW)},
		{Title: "Great tackler", Color: string(GREEN)},
		{Title: "Incredible passer", Color: string(GREEN)},
		{Title: "High AA", Color: string(GREEN)},
	}
}
