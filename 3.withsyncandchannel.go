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

/**
在这段代码中，通过 `ch` 通道和 `sync.WaitGroup` 来控制并发 Goroutines 数量。这里对 Goroutines 数量做了限制，但由于 `taskCnt` 被设置为 `math.MaxInt64`（极大的值），代码仍然可能导致内存耗尽或系统崩溃。以下是代码的详细解释：

### 代码分析

1. **全局变量 `wg`**：
   - `wg` 是一个 `sync.WaitGroup` 实例，用于等待所有 Goroutines 完成。

2. **`worker` 函数**：
   - 每个 `worker` 函数接收一个通道 `ch` 和任务编号 `i`，打印任务编号和当前 Goroutine 数量。
   - 从 `ch` 通道读取一个值后执行 `wg.Done()`，表示完成该任务并释放 `WaitGroup` 计数。

3. **`main` 函数**：
   - `taskCnt` 设置为 `math.MaxInt64`，表示一个非常大的任务数。
   - `ch` 通道的缓冲区大小为 3，这样最多允许 3 个 Goroutines 同时运行。
   - `for` 循环每次增加 `wg` 计数并向 `ch` 写入一个值，阻塞新 Goroutines 直到有缓冲区可用，然后启动一个 `worker` Goroutine。
   - `wg.Wait()` 阻塞主 Goroutine，直到所有 Goroutines 完成。

### 问题
虽然 `ch` 的容量限制了并发 Goroutines 数量，但 `taskCnt` 过大，`for` 循环将无限制地持续创建 Goroutines。由于 `wg.Add(1)` 每次都会增加计数，最终会超出可处理范围，导致内存泄漏和系统崩溃。

### 改进建议
应避免将 `taskCnt` 设为极大的值，或者通过 Goroutine 池的方式管理有限的任务。以下是一个改进版本，设定了一个合理的任务数：

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func worker(ch chan bool, i int) {
	defer wg.Done()
	fmt.Println("go func", i, " goroutine count = ", runtime.NumGoroutine())
	<-ch
}

func main() {
	taskCnt := 100     // 设置合理的任务数量
	ch := make(chan bool, 3)  // 控制并发的 Goroutines 数量

	for i := 0; i < taskCnt; i++ {
		wg.Add(1)
		ch <- true
		go worker(ch, i)
	}

	wg.Wait()
}
```

### 改进后的代码工作原理

1. 设置 `taskCnt` 为一个较小的合理值，例如 100，避免过多的任务创建。
2. 使用带缓冲的通道 `ch` 控制并发数量。通道容量设置为 3，这样最多有 3 个 Goroutines 可以同时运行。
3. `wg.Wait()` 等待所有 Goroutines 完成，确保所有任务完成后才退出主程序。
*/