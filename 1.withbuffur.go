package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	taskCnt := math.MaxInt64

	ch := make(chan bool, 3)

	for i := 0; i < taskCnt; i++ {
		ch <- true
		go worker(ch, i)
	}
}

func worker(ch chan bool, i int) {
	fmt.Println("go func", i, " goroutine count = ", runtime.NumGoroutine())
	<-ch
}
