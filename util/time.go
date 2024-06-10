package util

import "time"

func Parse(tStr string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05.000Z", tStr)
}

func TryParse(tStr string) time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05.000Z", tStr)
	return t
}

func MustParse(tStr string) time.Time {
	t, err := time.Parse("2006-01-02T15:04:05.000Z", tStr)
	if err != nil {
		panic(err)
	}
	return t
}
