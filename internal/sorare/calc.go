package sorare

import (
	"sync"
)

func GetAllPlayersNames() []PlayerName {
	var wg sync.WaitGroup
	leagues := getDomesticLeagueSlugs()
	clubsChan := make(chan string)
	wg.Add(len(leagues))
	for _, l := range leagues {
		go func(slug string) {
			defer wg.Done()
			clubs := getClubsSlugs(slug)
			for _, c := range clubs {
				clubsChan <- c
			}
		}(l)
	}
	go func() {
		wg.Wait()
		close(clubsChan)
	}()
	var wg2 sync.WaitGroup
	playersChan := make(chan PlayerName)
	for c := range clubsChan {
		go func(slug string) {
			wg2.Add(1)
			defer wg2.Done()
			players := getPlayersSlugs(slug)
			for _, p := range players {
				playersChan <- p
			}
		}(c)
	}
	go func() {
		wg2.Wait()
		close(playersChan)
	}()
	var ret []PlayerName
	for p := range playersChan {
		ret = append(ret, p)
	}
	return ret
}
