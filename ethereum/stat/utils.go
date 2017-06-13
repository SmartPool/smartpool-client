package stat

import (
	"time"
)

const (
	BaseTimePeriod int64 = 10 * 60
	ShortWindow    int64 = 2 * 6 * BaseTimePeriod
	LongWindow     int64 = 24 * 6 * BaseTimePeriod
)

var (
	Zone = LoadTimeZone()
)

func LoadTimeZone() *time.Location {
	loc, _ := time.LoadLocation("Local")
	return loc
}

func TimeToPeriod(t time.Time) uint64 {
	return uint64(t.Unix() / BaseTimePeriod)
}
