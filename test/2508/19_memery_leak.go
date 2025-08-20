package main

import (
	"fmt"
	"runtime"
	"time"
)

// memoryLeak function creates a goroutine that leaks.
func memoryLeak() {
	ch := make(chan int)
	go func() {
		// This goroutine will block forever because nothing is receiving from the channel.
		// The memory for the channel and the goroutine's stack will be leaked.
		ch <- 1
	}()
}

func main() {
	fmt.Println("Number of goroutines at start:", runtime.NumGoroutine())

	for i := 0; i < 10; i++ {
		memoryLeak()
	}

	// Give the goroutines time to start and block.
	time.Sleep(time.Second)

	fmt.Println("Number of goroutines after creating leaks:", runtime.NumGoroutine())
	// We expect to see 11 goroutines here (main + 10 leaked goroutines).

	// In a real application, the program would continue to do other work.
	// The leaked goroutines would continue to consume memory.
	// To demonstrate, we can force a garbage collection and check the memory stats.
	runtime.GC()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf("	TotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("	Sys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("	NumGC = %v ", m.NumGC)

	// Keep the main goroutine alive to observe the leak from outside the program (e.g., using pprof).
	// In this example, we'll just sleep for a bit.
	time.Sleep(10 * time.Second)

	fmt.Println("Number of goroutines at end:", runtime.NumGoroutine())
}
