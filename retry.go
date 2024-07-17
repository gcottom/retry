package retry

import (
	"errors"
	"fmt"
	"reflect"
)

// RetryAlgorithm is an interface that defines the methods that a retry algorithm must implement.
// SleepFunc is a method that will be called when the function should sleep before retrying.
// Reset is a method that will be called when the function should reset the algorithm.
// Clone is a method that will be called when the function should clone the algorithm.
type RetryAlgorithm interface {
	SleepFunc()
	Reset()
	Clone() RetryAlgorithm
}

// RetryWithLogger is a function that retries a function call using the provided retry algorithm.
// It takes the retry algorithm, the maximum number of retries, a logging function that takes an error as an argument,
// the function to call, and the arguments to pass to the function. It returns the results of the function call and an
// error if the function call fails. RetryWithLogger always returns the results of the function call as a slice of any, even if
// the function returns a single value. RetryWithLogger returns the exact number of results that the function retryableFn returns
// unless the function returns an error. In which case RetryWithLogger will retry the function call using the provided
// retry algorithm until the maximum number of retries is reached. If the function call fails after the maximum number
// of retries, RetryWithLogger will an empty slice of any and the error that caused the function call to fail.
func RetryWithLogger(alg RetryAlgorithm, maxRetries int, loggerFn func(err error), retryableFn any, args ...any) ([]any, error) {
	fnValue := reflect.ValueOf(retryableFn)
	if fnValue.Kind() != reflect.Func {
		return []any{}, errors.New("provided argument is not a function")
	}
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	var result []reflect.Value
	var err error

	for i := 0; i < maxRetries; i++ {
		result = fnValue.Call(in)
		errValue := result[len(result)-1]

		if !errValue.IsNil() {
			err = errValue.Interface().(error)
			loggerFn(err)
			alg.SleepFunc()
		} else {
			results := make([]interface{}, len(result)-1)
			for j := 0; j < len(result)-1; j++ {
				results[j] = result[j].Interface()
			}
			alg.Reset()
			return results, nil
		}
	}
	alg.Reset()
	return []any{}, err
}

// Retry is a function that retries a function call using the provided retry algorithm.
// It takes the retry algorithm, the maximum number of retries, the function to call,
// and the arguments to pass to the function. It returns the results of the function call and an
// error if the function call fails. Retry always returns the results of the function call as a slice of any, even if
// the function returns a single value. Retry returns the exact number of results that the function retryableFn returns
// unless the function returns an error. In which case Retry will retry the function call using the provided
// retry algorithm until the maximum number of retries is reached. If the function call fails after the maximum number
// of retries, Retry will an empty slice of any and the error that caused the function call to fail.
func Retry(alg RetryAlgorithm, maxRetries int, retryableFn any, args ...any) ([]any, error) {
	fnValue := reflect.ValueOf(retryableFn)
	if fnValue.Kind() != reflect.Func {
		return []any{}, errors.New("provided argument is not a function")
	}
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	var result []reflect.Value
	var err error

	for i := 0; i < maxRetries; i++ {
		result = fnValue.Call(in)
		errValue := result[len(result)-1]

		if !errValue.IsNil() {
			err = errValue.Interface().(error)
			fmt.Printf("error: %v, retrying...\n", err)
			alg.SleepFunc()
		} else {
			results := make([]interface{}, len(result)-1)
			for j := 0; j < len(result)-1; j++ {
				results[j] = result[j].Interface()
			}
			alg.Reset()
			return results, nil
		}
	}
	alg.Reset()
	return []any{}, err
}
