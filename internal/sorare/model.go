package sorare

type OpenLeagues struct {
	Football struct {
		LeaguesOpenForGameStats []struct {
			Format string `json:"format"`
			Slug   string `json:"slug"`
		} `json:"leaguesOpenForGameStats"`
	} `json:"football"`
}

type League struct {
	Football struct {
		Competition struct {
			Clubs struct {
				Nodes []struct {
					Slug string `json:"slug"`
				} `json:"nodes"`
			} `json:"clubs"`
		} `json:"competition"`
	} `json:"football"`
}

type Team struct {
	Football struct {
		Club struct {
			ActivePlayers struct {
				Nodes []struct {
					Slug        string `json:"slug"`
					DisplayName string `json:"displayName"`
				} `json:"nodes"`
			} `json:"activePlayers"`
		} `json:"club"`
	} `json:"football"`
}

type Player struct {
	Football struct {
		Player struct {
			Slug        string `json:"slug"`
			PictureUrl  string `json:"pictureUrl"`
			DisplayName string `json:"displayName"`
			Position    string `json:"position"`
			Age         int    `json:"age"`
			Country     struct {
				FlagUrl         string `json:"flagUrl"`
				ThreeLetterCode string `json:"threeLetterCode"`
			} `json:"country"`
			ActiveClub struct {
				PictureUrl string `json:"pictureUrl"`
				ShortName  string `json:"shortName"`
			} `json:"activeClub"`
			AllSo5Scores struct {
				Nodes []Score `json:"nodes"`
			} `json:"allSo5Scores"`
		} `json:"player"`
	} `json:"football"`
}

type Score struct {
	Score         float32 `json:"score"`
	DetailedScore []struct {
		Stat       string  `json:"stat"`
		StatValue  float32 `json:"statValue"`
		Category   string  `json:"category"`
		TotalScore float32 `json:"totalScore"`
	} `json:"detailedScore"`
	PlayerGameStats struct {
		MinsPlayed int `json:"minsPlayed"`
		Team       struct {
			Slug string `json:"slug"`
		} `json:"team"`
	} `json:"playerGameStats"`
	Game struct {
		Date      string `json:"date"`
		HomeGoals int    `json:"homeGoals"`
		AwayGoals int    `json:"awayGoals"`
		HomeTeam  struct {
			Slug       string `json:"slug"`
			PictureUrl string `json:"pictureUrl"`
		} `json:"homeTeam"`
		AwayTeam struct {
			Slug       string `json:"slug"`
			PictureUrl string `json:"pictureUrl"`
		} `json:"awayTeam"`
	} `json:"game"`
}

type PlayerName struct {
	DisplayName string
	Slug        string
}
