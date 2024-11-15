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

/**
这段代码的设计是通过一个任务通道 `ch` 将大量任务分配给固定数量的 Goroutines，但因为任务数量设置为 `math.MaxInt64`（极大值），即使使用了 Goroutines 池，也会导致系统资源耗尽或崩溃。以下是对代码的详细分析：

### 代码分析

1. **全局变量 `wg`**：
   - `wg` 是 `sync.WaitGroup` 实例，用于等待所有任务完成。

2. **`worker` 函数**：
   - `worker` 函数接受通道 `ch`，不断从 `ch` 中读取任务 `t`，并打印当前任务编号和 Goroutine 数量。
   - 每完成一个任务，调用 `wg.Done()` 减少计数。

3. **`sendTask` 函数**：
   - `sendTask` 函数将任务发送到通道 `ch`，并在每次发送前调用 `wg.Add(1)` 增加计数。

4. **`main` 函数**：
   - `ch` 是一个无缓冲通道，用于传递任务。
   - `goCnt` 设置为 3，表示会启动 3 个 Goroutines 执行 `worker` 函数，形成一个固定大小的 Goroutines 池。
   - `for t := 0; t < math.MaxInt64; t++` 启动了一个无限循环，为每个任务调用 `sendTask` 并将任务编号发送到通道 `ch`。
   - `wg.Wait()` 阻塞主 Goroutine，直到所有任务完成。

### 问题

1. **资源耗尽**：`math.MaxInt64` 是极大的数值，代码会尝试创建无限的任务并添加到 `wg` 中，很快会用尽内存资源。
2. **无限制任务生成**：即使有 Goroutine 池处理任务，这种无限制地生成任务和增加 `wg` 计数的方式无法真正有效地控制任务数量和资源使用。

### 改进建议

为了使程序更加稳定和实际，应该限制任务数量，并使用 Goroutine 池来平衡任务生成和消费。下面是一个更改后的版本，将任务数量设为合理值，并使用带缓冲的通道限制并发任务数量：

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func worker(ch chan int) {
	for t := range ch {
		fmt.Println("go task =", t, ", goroutine count =", runtime.NumGoroutine())
		wg.Done()
	}
}

func sendTask(task int, ch chan int) {
	wg.Add(1)
	ch <- task
}

func main() {
	taskCnt := 100         // 设定合理的任务总数
	ch := make(chan int, 10) // 使用带缓冲的通道以限制任务提交速度
	goCnt := 3               // 启动固定数量的 Goroutine

	for i := 0; i < goCnt; i++ {
		go worker(ch)
	}

	for t := 0; t < taskCnt; t++ {
		sendTask(t, ch)
	}

	wg.Wait()
	close(ch) // 所有任务完成后关闭通道，确保 `worker` Goroutines 退出
}
```

### 改进后代码的工作原理

1. 将 `taskCnt` 设置为有限值（例如 100），防止无限制任务创建。
2. 使用带缓冲的通道 `ch` 控制任务的发送速度。
3. 通过固定的 Goroutine 池 (`goCnt`) 来处理任务，限制同时运行的 Goroutines 数量。
4. 主 Goroutine 调用 `wg.Wait()`，等待所有任务完成后，关闭通道 `ch`，让 `worker` 函数的 `for` 循环退出。

这样就避免了无限制的资源消耗问题，同时达到了并发处理任务的效果。
*/