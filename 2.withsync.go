package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func worker(i int) {
	fmt.Println("go func", i, " goroutine count = ", runtime.NumGoroutine())
	wg.Done()
}

// don't work
func main() {
	taskCnt := math.MaxInt64

	for i := 0; i < taskCnt; i++ {
		wg.Add(1)
		go worker(1)
	}

	wg.Wait()
}
