package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomArray(n int, limit int) []int {
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(limit)
	}
	return arr
}

func testAlgorithm(sort func([]int) []int, size int, limit int, iterationNumber int) {
	for i := 0; i < iterationNumber; i++ {
		arr := generateRandomArray(size, limit)

		sort(arr)

		fmt.Println("Sort result:", arr[:5], arr[len(arr)-5:])
	}
}

func timeAlgorithm(sort func([]int) []int, size int, limit int, iterationNumber int) {
	var averageTime float64

	for i := 0; i < iterationNumber; i++ {
		arr := generateRandomArray(size, limit)

		start := time.Now()
		sort(arr)
		end := time.Since(start).Seconds()

		averageTime += end

		fmt.Printf("Sort with %d elements: %f seconds\n", size, end)
	}

	fmt.Printf("Average time: %f seconds\n", averageTime/float64(iterationNumber))
}
