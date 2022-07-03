package Futures

import (
	"time"
)

type Future[T any] struct {
	result *T
	done   chan bool
}

// Async a simple way to turn a function into a Future
// if the function you want to execute takes any parameters consider wrapping it in an anonymous function
func Async[T any](f func() T) Future[T] {
	var retVal T
	done := make(chan bool, 1)

	go func() {
		defer close(done)
		retVal = f()
		done <- true
	}()

	return Future[T]{
		result: &retVal,
		done:   done,
	}
}

// Await blocks until the Future is resolved
// returns the Future's underlying value
// may block indefinitely
func (f Future[T]) Await() T {
	<-f.done
	return *f.result
}

// AwaitWithTimeout blocks until the Future is resolved or the timeout was reached
// returns either a pointer to the Future's underlying (*value, true)
// or (nil, false) if the timeout was reached
// a timeout does NOT stop the Future's underlying go routine
func (f Future[T]) AwaitWithTimeout(timeout time.Duration) (*T, bool) {
	for {
		select {
		case <-time.After(timeout):
			return nil, false
		case <-f.done:
			return f.result, true
		}
	}
}
