package sorare

var playersInfoQuery string = `
query {
	football {
		player(slug:"%s") {
			slug
			displayName
			position
			age
			pictureUrl
			country {
				flagUrl
				threeLetterCode
			}
			activeClub {
				pictureUrl
				name
				shortName
			}
			allSo5Scores(first:50) {
				nodes {
					score
					detailedScore {
						stat
						statValue
						category
						totalScore
					}
					playerGameStats {
						minsPlayed 
						team {
							... on TeamInterface {
								slug
							}
						}
					}
					game {
						date
						homeGoals
						awayGoals
						homeTeam {
							... on TeamInterface {
								slug
								pictureUrl
							}
						}
						awayTeam {
							... on TeamInterface {
								slug
								pictureUrl
							}
						}
					}
				}
			}
		}
	}
}
`

var openLeaguesQuery = `
query {
	football {
		leaguesOpenForGameStats {
			format
			slug
		}
	}
}
`

var leagueQuery = `
query {
	football {
		competition(slug:"%s") {
			clubs {
				nodes {
					slug
				}
			}
		}
	}
}
`

var teamQuery = `
query {
	football {
		club(slug:"%s") {
			activePlayers {
				nodes {
					slug
					displayName
				}
			}
		}
	}
}
`
