package retry

import (
	"math"
	"math/rand/v2"
	"time"
)

// AlgExp is a retry algorithm that sleeps for an exponential interval.
// The exponential interval is calculated by multiplying the base by 2 to the power of the retry count.
type AlgExp struct {
	Base         int
	RetryCount   int
	TimeDuration time.Duration
}

// AlgExpJitter is a retry algorithm that sleeps for an exponential interval with jitter.
// The exponential interval is calculated by multiplying the base by 2 to the power of the retry count multiplied by a random float64 + 0.5.
type AlgExpJitter struct {
	Base         int
	RetryCount   int
	TimeDuration time.Duration
}

// AlgExp is a retry algorithm that sleeps for an exponential interval.
// NewAlgExpDefault creates a new AlgExp with a default base of 1000 and a default time duration of time.Millisecond.
func NewAlgExpDefault() *AlgExp {
	return &AlgExp{Base: 1000, RetryCount: 1, TimeDuration: time.Millisecond}
}

// AlgExp is a retry algorithm that sleeps for an exponential interval.
// NewAlgExp creates a new AlgExp with the given base and time duration.
func NewAlgExp(base int, duration time.Duration) *AlgExp {
	if base < 0 {
		base = 0
	}
	return &AlgExp{Base: base, RetryCount: 1, TimeDuration: duration}
}

// AlgExpJitter is a retry algorithm that sleeps for an exponential interval with jitter.
// NewAlgExpJitterDefault creates a new AlgExpJitter with a default base of 1000 and a default time duration of time.Millisecond.
func NewAlgExpJitterDefault() *AlgExpJitter {
	return &AlgExpJitter{Base: 1000, RetryCount: 1, TimeDuration: time.Millisecond}
}

// AlgExpJitter is a retry algorithm that sleeps for an exponential interval with jitter.
// NewAlgExpJitter creates a new AlgExpJitter with the given base and time duration.
func NewAlgExpJitter(base int, duration time.Duration) *AlgExpJitter {
	if base < 0 {
		base = 0
	}
	return &AlgExpJitter{Base: base, RetryCount: 1, TimeDuration: duration}
}

func (a *AlgExp) SleepFunc() {
	time.Sleep(time.Duration(a.Base) * a.TimeDuration * time.Duration(1<<uint(a.RetryCount)))
	a.RetryCount++
}

func (a *AlgExpJitter) SleepFunc() {
	time.Sleep(time.Duration(float64(a.Base)*(math.Pow(2, float64(a.RetryCount))*(rand.Float64()+0.5))) * a.TimeDuration)
	a.RetryCount++
}

func (a *AlgExp) Reset() {
	a.RetryCount = 1
}

func (a *AlgExpJitter) Reset() {
	a.RetryCount = 1
}

func (a *AlgExp) Clone() RetryAlgorithm {
	return NewAlgExp(a.Base, a.TimeDuration)
}

func (a *AlgExpJitter) Clone() RetryAlgorithm {
	return NewAlgExpJitter(a.Base, a.TimeDuration)
}
