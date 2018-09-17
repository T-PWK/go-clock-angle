package clock

import (
	"math"
)

const (
	minHour   = 0
	maxHour   = 24
	minMinSec = 0
	maxMinSec = 59

	fullAngle = 360

	hourAngle        = fullAngle / 12.0
	minuteAngle      = fullAngle / 60.0
	minuteDeltaAngle = minuteAngle / 60
	hourDeltaAngle   = hourAngle / 60.0
	overlapConst     = 30.0 / 5.5
)

// Time struct representing time
type Time struct {
	Hour, Min, Sec int
}

// Overlaps returns a slice of times where the minute hand overlaps hour hand
func Overlaps() []Time {
	times := make([]Time, 0, 12)

	for h := 0; h < 12; h++ {
		m, s := math.Modf(overlapConst * float64(h))
		min := int(m)
		sec := int(math.Round(s * 60))

		if sec == 60 {
			continue
		}

		times = append(times, Time{h, min, sec})
	}

	return times
}

// AngleMinutesToHours returns angle from the minute to the hour hand
func AngleMinutesToHours(h, m, s int) int {
	if h < minHour || h > maxHour || m < minMinSec || m > maxMinSec || s < minMinSec || s > maxMinSec {
		panic("Invalid hour, minute or second value")
	}

	h = h % 12

	ha := float64(h)*hourAngle + hourDeltaAngle*float64(m)
	ma := float64(m)*minuteAngle + minuteDeltaAngle*float64(s)

	a := int(ha - ma)

	if a < 0 {
		a = fullAngle + a
	}

	return a
}

// AngleHoursToMinutes returns angle from the hour to the minute hand
func AngleHoursToMinutes(h, m, s int) int {
	return (fullAngle - AngleMinutesToHours(h, m, s)) % fullAngle
}
