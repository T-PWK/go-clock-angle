package clock

import (
	"testing"
)

func BenchmarkOverlaps(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Overlaps()
	}
}
func BenchmarkAngleMinutesToHours(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		AngleMinutesToHours(1, 50, 0)
	}
}
func BenchmarkAngleHoursToMinutes(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		AngleHoursToMinutes(1, 50, 0)
	}
}

func TestOverlaps(t *testing.T) {
	times := Overlaps()

	if len(times) != 11 {
		t.Errorf("Overlaps should return 11 records, got: %d", len(times))
	}
}

func TestInvalidInput(t *testing.T) {

	f := func(f func(int, int, int) int, h, m int) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Function should panic for hours: %d, minutes: %d", h, m)
			}
		}()

		f(h, m, 0)
	}

	f(AngleMinutesToHours, -1, 30)
	f(AngleMinutesToHours, 25, 30)
	f(AngleMinutesToHours, 10, -1)
	f(AngleMinutesToHours, 10, 60)

	f(AngleHoursToMinutes, -1, 30)
	f(AngleHoursToMinutes, 25, 30)
	f(AngleHoursToMinutes, 10, -1)
	f(AngleHoursToMinutes, 10, 60)
}

func TestAngles(t *testing.T) {
	tests := []struct {
		h, m, r1, r2 int
	}{
		{0, 0, 0, 0},
		{12, 0, 0, 0},
		{1, 50, 115, 245},
		{13, 50, 115, 245},
		{3, 0, 90, 270},
		{15, 0, 90, 270},
		{6, 0, 180, 180},
		{18, 0, 180, 180},
		{9, 0, 270, 90},
		{21, 0, 270, 90},
	}

	for _, test := range tests {
		if r := AngleMinutesToHours(test.h, test.m, 0); r != test.r1 {
			t.Errorf("AngleMinutesToHours failed, got: %d, want: %d", r, test.r1)
		}

		if r := AngleHoursToMinutes(test.h, test.m, 0); r != test.r2 {
			t.Errorf("AngleHoursToMinutes failed, got: %d, want: %d", r, test.r2)
		}
	}
}
