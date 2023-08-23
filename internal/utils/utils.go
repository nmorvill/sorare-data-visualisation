package utils

import "time"

func GetGameweekFromString(str string) int {
	date, err := time.Parse("2006-01-02T15:04:05Z", str)
	if err != nil {
		return 0
	}
	gw374 := time.Date(2023, time.May, 23, 14, 0, 0, 0, time.Local)
	daysDiff := -int(gw374.Sub(date).Hours() / 24)
	gwDiff := (daysDiff / 7) * 2
	if daysDiff%7 >= 3 {
		gwDiff += 1
	}
	return gwDiff + 374
}

func GetColorCodeOfNote(note int) string {
	if note >= 70 {
		return "#34732F"
	} else if note >= 60 {
		return "#5EA258"
	} else if note >= 50 {
		return "#8EBC8A"
	} else if note >= 40 {
		return "#F1AF4B"
	} else if note >= 30 {
		return "#CD6D15"
	} else if note >= 20 {
		return "#CD4115"
	}
	return "#CD4115"
}
