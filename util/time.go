package util

import "time"

func Parse(tStr string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05.000Z", tStr)
}

func TryParse(tStr string) time.Time {
	if t, err := time.Parse("2006-01-02T15:04:05.000Z", tStr); err != nil {
		t, _ := time.Parse("2006-01-02T15:04:05.000Z", tStr)
		return t
	} else {
		return t
	}
}

func MustParse(tStr string) time.Time {
	if t, err := time.Parse("2006-01-02T15:04:05.000Z", tStr); err != nil {
		t, err := time.Parse("2006-01-02T15:04:05.000Z", tStr)
		if err != nil {
			panic(err)
		} else {
			return t
		}
	} else {
		return t
	}
}

func Format(t *time.Time) string {
	return t.Format("2006-01-02T15:04:05.000Z")
}

func FormatV0(t *time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
