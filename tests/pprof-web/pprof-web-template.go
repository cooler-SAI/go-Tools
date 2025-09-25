package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

// INSERT YOUR TEST FUNCTION HERE
func yourFunctionToTest() {
	// Example: replace this code with your function
	fmt.Println("Executing your function...")
	time.Sleep(100 * time.Millisecond)

	// Memory test - creating a slice
	data := make([]byte, 1024*1024) // 1MB
	for i := range data {
		data[i] = byte(i % 256)
	}
}

// INSERT YOUR CPU-INTENSIVE FUNCTION HERE
func yourCPUIntensiveFunction() {
	// Example: replace with your CPU-intensive function
	sum := 0
	for i := 0; i < 1000000; i++ {
		sum += i * i
	}
}

// INSERT YOUR MEMORY-INTENSIVE FUNCTION HERE
func yourMemoryIntensiveFunction() {
	// Example: replace with your memory-intensive function
	var slices [][]byte
	for i := 0; i < 100; i++ {
		slice := make([]byte, 1024*1024) // 1MB
		slices = append(slices, slice)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	// HTTP handlers setup
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// CALL YOUR TEST FUNCTION HERE
		yourFunctionToTest()

		duration := time.Since(start)
		fprintf, err := fmt.Fprintf(w, "Function executed in %v\n", duration)
		if err != nil {
			fmt.Println(fprintf)
			return
		}
	})

	http.HandleFunc("/cpu-test", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// CPU TEST - INSERT YOUR CODE HERE
		for i := 0; i < 100; i++ {
			yourCPUIntensiveFunction()
		}

		duration := time.Since(start)
		fprintf, err := fmt.Fprintf(w, "CPU test completed in %v\n", duration)
		if err != nil {
			fmt.Println(fprintf)
			return
		}
	})

	http.HandleFunc("/memory-test", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// MEMORY TEST - INSERT YOUR CODE HERE
		yourMemoryIntensiveFunction()

		duration := time.Since(start)
		fprintf, err := fmt.Fprintf(w, "Memory test completed in %v\n", duration)
		if err != nil {
			fmt.Println(fprintf)
			return
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fprintf, err := fmt.Fprintf(w, `
				<h1>PProf Testing Dashboard</h1>
				<p>Use this to test your code with pprof profiling</p>
				<ul>
					<li><a href="/test">Test Main Function</a></li>
					<li><a href="/cpu-test">Test CPU Load</a></li>
					<li><a href="/memory-test">Test Memory Load</a></li>
					<li><a href="/debug/pprof/">PProf Profiling</a></li>
				</ul>
				
				<h3>How to use:</h3>
				<ol>
					<li>Replace the functions yourFunctionToTest, yourCPUIntensiveFunction, yourMemoryIntensiveFunction</li>
					<li>Run the application: go run pprof-web-template.go</li>
					<li>Open http://localhost:8080/debug/pprof/ for profiling</li>
					<li>Test the endpoints and analyze the profiles</li>
				</ol>
				`)
		if err != nil {
			fmt.Println(fprintf)
			return
		}
	})

	fmt.Println("Server started at http://localhost:8080")
	fmt.Println("PProf available at http://localhost:8080/debug/pprof/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
