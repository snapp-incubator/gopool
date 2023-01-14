# Gopool
A Golang goroutine pool with graceful shutdown.

## Installation

To install this package, you need to setup your Go workspace. The simplest way to install the library is to run:
```
$ go get github.com/snapp-incubator/gopool
```

## Example
```go
package main

import (
	"fmt"
	"time"

	"github.com/snapp-incubator/gopool"
)

type sample struct {
	counter int
}

func newSample() gopool.Job {
	return &sample{
		counter: 0,
	}
}

func (t *sample) Do() {
	fmt.Println("Task started at:", time.Now().Format("2006-01-02 15:04:05"))
	t.counter += 1
	fmt.Println("Task finished at:", time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
	wp := gopool.NewWorkerPool(2, 3)
	requests := []string{"alpha", "beta", "gamma", "delta", "epsilon"}

	for _, r := range requests {
		wp.AddTask(newSample())
		fmt.Printf("Handling Task: %s\n", r)
		time.Sleep(time.Second)
	}

	wp.Shutdown()
}

```