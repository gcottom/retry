package retry

// RetryManager is a manager for retrying functions with a given algorithm and max retries.
type RetryManager struct {
	Algorithm  RetryAlgorithm
	MaxRetries int
}

// NewRetryManager creates a new RetryManager with the given algorithm and max retries.
// RetryManager is safe for concurrent use as the Retry function on the RetryManager
// will clone the algorithm before using it. Ensuring that the algorithm is not modified
// by multiple goroutines.
func NewRetryManager(algorithm RetryAlgorithm, maxRetries int) *RetryManager {
	return &RetryManager{Algorithm: algorithm, MaxRetries: maxRetries}
}

// Retry is a function that retries a function call using the provided retry algorithm.
// It takes the function to call and the arguments to pass to the function. It returns the results of the function call and an
// error if the function call fails. Retry always returns the results of the function call as a slice of any, even if
// the function returns a single value. Retry returns the exact number of results that the function retryableFn returns
// unless the function returns an error. In which case Retry will retry the function call using the provided
// retry algorithm until the maximum number of retries is reached. If the function call fails after the maximum number
// of retries, Retry will an empty slice of any and the error that caused the function call to fail.
func (r *RetryManager) Retry(fn any, args ...any) ([]any, error) {
	alg := r.Algorithm.Clone()
	return Retry(alg, r.MaxRetries, fn, args...)
}

// RetryWithLogger is a function that retries a function call using the provided retry algorithm.
// It takes a logging function that takes an error as an argument,
// the function to call, and the arguments to pass to the function. It returns the results of the function call and an
// error if the function call fails. RetryWithLogger always returns the results of the function call as a slice of any, even if
// the function returns a single value. RetryWithLogger returns the exact number of results that the function retryableFn returns
// unless the function returns an error. In which case RetryWithLogger will retry the function call using the provided
// retry algorithm until the maximum number of retries is reached. If the function call fails after the maximum number
// of retries, RetryWithLogger will an empty slice of any and the error that caused the function call to fail.
func (r *RetryManager) RetryWithLogger(loggerFn func(err error), fn any, args ...any) ([]any, error) {
	alg := r.Algorithm.Clone()
	return RetryWithLogger(alg, r.MaxRetries, loggerFn, fn, args...)
}
