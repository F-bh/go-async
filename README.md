# go-async
A dead simple type safe implementation of the async - await pattern in golang

## But why?
### because:
1. it can be done
2. it cut's down on channel boilerplate 
3. it gives programmers coming from languages like Rust, TypeScript and C# a familiar interface for concurrent programming

## How to install:
- `go get https://github.com/F-bh/go-async`
- it requires at least golang version 1.18 to use 

## What is included
- 1 type called "Future"
- 1 function to turn a given function into a "Future"
   - Async()
- 2 methods to await the result of a given "Future"
    - Await()
    - AwaitWithTimeout()

## Examples
```
package examples

import (
	"github.com/F-bh/go-async"
	"time"
)

func example(){
	future := Futures.Async(func() string{
		time.Sleep(500)
		return "Hello World"
	})

	println(future.Await())

	futureTimeout := Futures.Async(func() string{
		time.Sleep(50000000)
		return "unused"
	})

	val, ok := futureTimeout.AwaitWithTimeout(500)
	if !ok{
		println("timed out")
	}else{
		println(*val)
	}
}
```
