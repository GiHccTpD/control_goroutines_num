package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func worker(ch chan bool, i int) {
	fmt.Println("go func", i, " goroutine count = ", runtime.NumGoroutine())
	<-ch

	wg.Done()
}

func main() {
	taskCnt := math.MaxInt64
	ch := make(chan bool, 3)

	for i := 0; i < taskCnt; i++ {
		wg.Add(1)
		ch <- true

		go worker(ch, i)
	}

	wg.Wait()
}
