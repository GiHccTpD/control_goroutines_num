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

/**
这段代码片段中，`taskCnt` 设置为 `math.MaxInt64`，意图创建非常多的 Goroutines，然而这种方式会导致程序资源耗尽并崩溃。以下是代码的详细分析：

### 代码分析

1. **导入包**：
   - `fmt`: 用于格式化输出。
   - `math`: 包含数学常量和函数，`math.MaxInt64` 是 `int64` 的最大值。
   - `runtime`: 提供与 Go 运行时环境的交互，`runtime.NumGoroutine()` 返回当前 Goroutine 数量。
   - `sync`: 包含 `sync.WaitGroup` 用于 Goroutine 的并发控制和同步。

2. **全局变量 `wg`**：
   - `wg` 是一个 `sync.WaitGroup`，用于等待所有 Goroutines 完成后再继续执行 `main` 函数的退出。

3. **worker 函数**：
   - 每个 `worker` 函数接收一个任务编号 `i`，打印 Goroutine 数量和任务编号 `i`，然后调用 `wg.Done()`，表示当前任务已完成，减少 `WaitGroup` 的计数。

4. **main 函数**：
   - `taskCnt` 设置为 `math.MaxInt64`，表示极大量的任务。
   - 使用 `for` 循环，每次调用 `wg.Add(1)` 增加 `WaitGroup` 计数，然后启动一个 `worker` Goroutine。
   - `wg.Wait()` 阻塞主 Goroutine，直到所有启动的 `worker` Goroutines 执行完毕并调用 `wg.Done()`。

### 问题
1. **资源耗尽**：因为 `taskCnt` 是一个极大的值，`for` 循环会尝试生成 `math.MaxInt64` 个 Goroutines。这样会很快耗尽系统资源，导致程序崩溃或死锁。
2. **不实际的任务量**：如此多的 Goroutines 即使运行，也会因资源限制无法按预期工作。

### 改进建议
可以限制并发 Goroutines 的数量，而不是尝试创建 `math.MaxInt64` 个 Goroutines。例如：

- **使用带缓冲的通道**：可以限制同时并发的 Goroutines 数量。
- **设定合理的任务数量**：将 `taskCnt` 设置为实际能运行的较小值，例如 100 或 1000。

### 示例改进代码

将 `taskCnt` 设置为一个较小值，并使用带缓冲的通道来限制并发：

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func worker(i int) {
	defer wg.Done()
	fmt.Println("go func", i, " goroutine count = ", runtime.NumGoroutine())
}

func main() {
	taskCnt := 100  // 设定一个合理的任务数量
	maxGoroutines := 5
	guard := make(chan struct{}, maxGoroutines)

	for i := 0; i < taskCnt; i++ {
		wg.Add(1)
		guard <- struct{}{} // 当缓冲区满时阻塞，限制 Goroutine 数量
		go func(i int) {
			defer func() { <-guard }() // 任务完成后释放缓冲区
			worker(i)
		}(i)
	}

	wg.Wait()
}
```

在这个改进的代码中，`guard` 通道的缓冲区大小控制了并发 Goroutines 的数量，确保在任何时间点最多只有 `maxGoroutines` 个 Goroutines 在运行。
*/