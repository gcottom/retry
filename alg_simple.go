package retry

import "time"

// AlgSimple is a simple retry algorithm that sleeps for a fixed interval.
type AlgSimple struct {
	Interval     int
	TimeDuration time.Duration
}

// AlgSimple is a simple retry algorithm that sleeps for a fixed interval.
// NewAlgSimpleDefault creates a new AlgSimple with a default interval of 1000 and a default time duration of time.Millisecond.
func NewAlgSimpleDefault() *AlgSimple {
	return &AlgSimple{Interval: 1000, TimeDuration: time.Millisecond}
}

// AlgSimple is a simple retry algorithm that sleeps for a fixed interval.
// NewAlgSimple creates a new AlgSimple with the given interval and time duration.
func NewAlgSimple(interval int, duration time.Duration) *AlgSimple {
	if interval < 0 {
		interval = 0
	}
	return &AlgSimple{Interval: interval, TimeDuration: duration}
}

func (a *AlgSimple) SleepFunc() {
	time.Sleep(time.Duration(a.Interval) * a.TimeDuration)
}

func (a *AlgSimple) Reset() {
	// no-op
}

func (a *AlgSimple) Clone() RetryAlgorithm {
	return NewAlgSimple(a.Interval, a.TimeDuration)
}
