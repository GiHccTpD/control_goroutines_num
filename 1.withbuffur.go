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

/**
这个代码片段尝试通过带缓冲的通道 `ch` 来限制并发 Goroutine 的数量，从而避免无限制地启动 Goroutines。以下是代码的详细解释：

1. **导入包**：
   - `fmt`: 用于格式化输出。
   - `math`: 包含数学常量和函数，`math.MaxInt64` 是 `int64` 类型的最大值。
   - `runtime`: 提供与 Go 运行时环境的交互，`runtime.NumGoroutine()` 返回当前 Goroutine 数量。

2. **main 函数**：
   - `taskCnt` 被定义为 `math.MaxInt64`，即 `int64` 的最大值，表示理论上的任务数量。
   - `ch` 是一个带缓冲的通道，容量为 3。该通道用于限制同时运行的 Goroutines 数量，确保在任何时候最多只有 3 个 Goroutines 在运行。
   - `for` 循环启动 `taskCnt` 次迭代。每次迭代时，将 `true` 写入通道 `ch`，然后启动一个新的 Goroutine `worker`。

3. **worker 函数**：
   - 每个 `worker` Goroutine 在启动时会打印消息，包含当前任务编号 `i` 和当前的 Goroutine 数量（通过 `runtime.NumGoroutine()` 获取）。
   - 执行完打印任务后，`worker` 会从通道 `ch` 读取并释放一个值，从而允许新的 Goroutine 启动。
   - 这种方式有效地控制了并发数量，因为当通道缓冲区已满（即 3 个值时），新的 Goroutine 会阻塞，直到缓冲区中有可用空间。

### 工作机制分析
- `ch` 的容量设置为 3，意味着最多只有 3 个 Goroutines 可以并发运行。每次循环写入 `ch` 都会占用一个位置，直到 `worker` 函数完成后从 `ch` 读取以释放位置。
- 这种模式类似于一个简单的 Goroutine 池，通道的缓冲区大小控制并发上限。

### 注意事项
即使通过 `ch` 控制并发数量，这段代码在大任务数下仍然存在问题：
   - `taskCnt` 的设置过高。理论上会无限生成任务，而 Goroutine 数量受限，但循环不会结束，因此会一直阻塞。
   - `taskCnt` 应设置为合理的数值，或使用外部条件来中断循环。
*/