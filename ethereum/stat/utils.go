package stat

import (
	"time"
)

const (
	BaseTimePeriod int64 = 60
	ShortWindow    int64 = 60 * BaseTimePeriod
	LongWindow     int64 = 3 * 60 * BaseTimePeriod
)

func TimeToPeriod(t time.Time) uint64 {
	return uint64(t.Unix() / BaseTimePeriod)
}
