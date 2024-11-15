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

/**
这个代码片段展示了一个无限制地启动 Goroutines 的例子。以下是代码的详细解释：

导入包：

fmt: 用于格式化输出。
math: 包含数学常量和函数，其中 math.MaxInt64 是 int64 类型的最大值。
runtime: 提供与 Go 运行时环境的交互，runtime.NumGoroutine() 可以返回当前正在运行的 Goroutine 数量。
main 函数：

定义 taskCnt 为 math.MaxInt64，即 int64 的最大值，相当于无穷大的任务数量。这样做实际上不太合理，因为无法启动如此多的 Goroutines，而且会很快耗尽系统资源。
for 循环使用 taskCnt，试图创建等同于 taskCnt 数量的 Goroutines，这在实践中是不可能的，因为会导致内存耗尽和系统崩溃。
启动 Goroutines：

每次循环，启动一个匿名 Goroutine，并打印一条消息。每个 Goroutine 会接收循环变量 i 的当前值并打印出来，同时调用 runtime.NumGoroutine() 打印当前 Goroutine 的数量。
注意这里的匿名函数参数 i int，将 i 传递到 Goroutine 内部，避免了共享变量 i 的问题（如果没有传递参数 i，会导致闭包捕获的 i 值可能不一致）。
问题分析
由于 taskCnt 被设置为一个极大的值，这样的代码会：

不断地创建 Goroutines，耗尽内存和 CPU。
导致系统资源耗尽，程序很快崩溃。
建议改进
可以限制 Goroutines 数量（如使用 Goroutine 池）或者减少 taskCnt 的数量。
*/