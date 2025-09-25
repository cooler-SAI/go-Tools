package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

// INSERT YOUR TEST FUNCTION HERE
func yourFunctionToTest() {
	// REPLACE THIS CODE WITH YOUR FUNCTION
	fmt.Println("Executing your test function...")

	// Example: complex calculations
	result := 0
	for i := 0; i < 1000000; i++ {
		result += i * i
	}

	// Example: memory operations
	data := make([]string, 1000)
	for i := range data {
		data[i] = fmt.Sprintf("string%d", i)
	}
}

// INSERT YOUR CPU-INTENSIVE FUNCTION HERE
func yourCPUIntensiveFunction() {
	// REPLACE WITH YOUR CPU-INTENSIVE FUNCTION
	total := 0
	for i := 0; i < 500000; i++ {
		for j := 0; j < 1000; j++ {
			total += i * j
		}
	}
}

// INSERT YOUR MEMORY-INTENSIVE FUNCTION HERE
func yourMemoryIntensiveFunction() {
	// REPLACE WITH YOUR MEMORY-INTENSIVE FUNCTION
	var collection [][]byte

	for i := 0; i < 50; i++ {
		// Allocate 1MB of memory
		chunk := make([]byte, 1024*1024)
		for j := range chunk {
			chunk[j] = byte((i + j) % 256)
		}
		collection = append(collection, chunk)
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	fmt.Println("=== PProf Local Testing Bench ===")
	fmt.Println("Replace the functions in the code with your own for testing")

	// CREATE CPU PROFILE
	cpuProfile, err := os.Create("cpu.pprof")
	if err != nil {
		log.Fatal("Error creating CPU profile:", err)
	}
	defer func(cpuProfile *os.File) {
		err := cpuProfile.Close()
		if err != nil {
			log.Fatal("Error closing CPU profile:", err)
		}
	}(cpuProfile)

	// Start CPU profiling
	if err := pprof.StartCPUProfile(cpuProfile); err != nil {
		log.Fatal("Error starting CPU profiling:", err)
	}

	fmt.Println("\n1. TESTING CPU LOAD...")

	// EXECUTE YOUR CODE FOR CPU PROFILING
	start := time.Now()

	// INSERT YOUR CODE FOR CPU TESTING HERE
	for i := 0; i < 10; i++ {
		yourCPUIntensiveFunction()
		fmt.Printf("CPU iteration %d completed\n", i+1)
	}

	cpuDuration := time.Since(start)
	pprof.StopCPUProfile()
	fmt.Printf("CPU profiling completed in %v\n", cpuDuration)

	fmt.Println("\n2. TESTING MEMORY...")

	// EXECUTE YOUR CODE FOR MEMORY PROFILING
	yourMemoryIntensiveFunction()

	fmt.Println("\n3. TESTING MAIN FUNCTION...")

	// TESTING YOUR MAIN FUNCTION
	for i := 0; i < 5; i++ {
		yourFunctionToTest()
		fmt.Printf("Function test %d completed\n", i+1)
	}

	// CREATE MEMORY PROFILE
	memoryProfile, err := os.Create("memory.pprof")
	if err != nil {
		log.Fatal("Error creating memory profile:", err)
	}
	defer func(memoryProfile *os.File) {
		err := memoryProfile.Close()
		if err != nil {
			log.Fatal("Error closing memory profile:", err)
		}
	}(memoryProfile)

	// Force garbage collection before creating memory profile
	runtime.GC()

	if err := pprof.WriteHeapProfile(memoryProfile); err != nil {
		log.Fatal("Error writing memory profile:", err)
	}

	// CREATE GOROUTINE PROFILE
	goroutineProfile, err := os.Create("goroutine.pprof")
	if err != nil {
		log.Fatal("Error creating goroutine profile:", err)
	}
	defer func(goroutineProfile *os.File) {
		err := goroutineProfile.Close()
		if err != nil {
			log.Fatal("Error closing goroutine profile:", err)
		}
	}(goroutineProfile)

	if goroutinePprof := pprof.Lookup("goroutine"); goroutinePprof != nil {
		err := goroutinePprof.WriteTo(goroutineProfile, 0)
		if err != nil {
			log.Fatal("Error writing goroutine profile:", err)
			return
		}
	}

	fmt.Println("\n=== PROFILES SAVED ===")
	fmt.Println("cpu.pprof - CPU profile")
	fmt.Println("memory.pprof - memory profile")
	fmt.Println("goroutine.pprof - goroutine profile")

	fmt.Println("\n=== HOW TO ANALYZE ===")
	fmt.Println("go tool pprof cpu.pprof")
	fmt.Println("go tool pprof memory.pprof")
	fmt.Println("(web) - for browser visualization")
}
