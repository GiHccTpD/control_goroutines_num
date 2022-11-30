package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func worker(ch chan int) {
	for t := range ch {
		fmt.Println("go task = ", t, ",goroutine count = ", runtime.NumGoroutine())
		wg.Done()
	}
}

func sendTask(task int, ch chan int) {
	wg.Add(1)
	ch <- task
}

func main() {
	ch := make(chan int)

	goCnt := 3 // 启动goroutine的数量

	for i := 0; i < goCnt; i++ {
		// 启动 处理任务
		go worker(ch)
	}

	for t := 0; t < math.MaxInt64; t++ {
		// 发送任务
		go sendTask(t, ch)
	}

	wg.Wait()
}
