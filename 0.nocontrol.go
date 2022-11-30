package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	taskCnt := math.MaxInt64

	for i := 0; i < taskCnt; i++ {
		go func(i int) {
			fmt.Println("go func", i, " goroutine count = ", runtime.NumGoroutine())
		}(i)
	}
}
