package stat

import (
	"time"
)

var BaseTimePeriod int64 = 60

func TimeToPeriod(t time.Time) uint64 {
	return uint64(t.Unix() / BaseTimePeriod)
}
