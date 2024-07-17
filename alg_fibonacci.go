package retry

import "time"

// AlgFibonacci is a retry algorithm that sleeps for a fibonacci interval.
// The fibonacci interval is calculated by adding the previous two values in the sequence.
type AlgFibonacci struct {
	Start1       int
	Start2       int
	Val1         int
	Val2         int
	TimeDuration time.Duration
}

// AlgFibonacci is a retry algorithm that sleeps for a fibonacci interval.
// NewAlgFibonacciDefault creates a new AlgFibonacci with a default start1 of 0, start2 of 1, and a default time duration of time.Millisecond.
func NewAlgFibonacciDefault() *AlgFibonacci {
	return &AlgFibonacci{Start1: 0, Start2: 1, Val1: 0, Val2: 1, TimeDuration: time.Millisecond}
}

// AlgFibonacci is a retry algorithm that sleeps for a fibonacci interval.
// NewAlgFibonacci creates a new AlgFibonacci with the given start1, start2, and time duration.
func NewAlgFibonacci(start1 int, start2 int, duration time.Duration) *AlgFibonacci {
	if start1 < 0 {
		start1 = 0
	}
	if start2 < 1 {
		start2 = 1
	}
	return &AlgFibonacci{Start1: start1, Start2: start2, Val1: start1, Val2: start2, TimeDuration: duration}
}

func (a *AlgFibonacci) SleepFunc() {
	time.Sleep(time.Duration(a.Val2) * 1000 * a.TimeDuration)
	a.Val1, a.Val2 = a.Val2, a.Val1+a.Val2
}

func (a *AlgFibonacci) Reset() {
	a.Val1, a.Val2 = a.Start1, a.Start2
}

func (a *AlgFibonacci) Clone() RetryAlgorithm {
	return NewAlgFibonacci(a.Start1, a.Start2, a.TimeDuration)
}
