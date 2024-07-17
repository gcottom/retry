# Retry

Retry is a simple generic function retryer. It comes with 4 predefined retry timing algorithms and the flexibility to add and define your own.

## Retry Comes In 2 Flavors
1. A simple function that takes a RetryAlgorithm, max number of retries, the function name that you want to perform retries on, and its arguments.
2. A RetryManager which holds your retry config and is safe for concurrent use. The RetryManager has a Retry function on it which is a wrapper for the
retry function in number 1. This wrapper uses the config in the RetryManager. The Retry function on the RetryManager only requires the function name
that is to be called and its list of arguments.

Both versions return a slice of any, and an error. The function will retry your function up to the max retry value specified. If on the last try,
your function returns an error, Retry will return an empty slice of any and the error from the function. Both versions also come in a WithLogger version. 
The WithLogger versions take a logger function that takes an error as an argument. 

## License
This project is licensed under the MIT License. See the LICENSE file for details.
