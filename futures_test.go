package Futures

import (
	"testing"
	"time"
)

func TestSimpleAwaitSuccess(t *testing.T) {
	f := Async(func() string {
		time.Sleep(time.Second * 1)
		return "success"
	})

	res := f.Await()

	if res != "success" {
		t.FailNow()
	}
}

func TestAwaitTimeout(t *testing.T) {
	f := Async(func() string {
		time.Sleep(time.Second * 3)
		return "failure"
	})

	res, ok := f.AwaitWithTimeout(time.Second * 1)
	if ok {
		t.FailNow()
	}
	if res != nil && *res == "failure" {
		t.FailNow()
	}
}

func TestAwaitNoTimeout(t *testing.T) {
	f := Async(func() string {
		time.Sleep(time.Second * 1)
		return "success"
	})

	res, ok := f.AwaitWithTimeout(time.Second * 3)
	if !ok {
		t.FailNow()
	}
	if res == nil || *res != "success" {
		t.FailNow()
	}
}
