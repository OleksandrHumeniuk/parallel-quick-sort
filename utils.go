package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomArray(n int, limit int) []interface{} {
	arr := make([]interface{}, n)

	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(limit)
	}
	return arr
}

func checkSorted(arr []interface{}, comp Comparator) bool {
	for i := 0; i < len(arr)-1; i++ {
		if comp(arr[i+1], arr[i]) {
			return false
		}
	}
	return true
}

func testAlgorithm(
	sort func([]interface{}, Comparator) []interface{},
	size int,
	limit int,
	iterationNumber int,
	comp Comparator) {
	for i := 0; i < iterationNumber; i++ {
		arr := generateRandomArray(size, limit)

		sort(arr, comp)

		fmt.Println("Array is sorted:", checkSorted(arr, comp))
		fmt.Println("Sort result:", arr[:5], arr[len(arr)-5:])
	}
}

func timeAlgorithm(
	sort func([]interface{}, Comparator) []interface{},
	size int,
	limit int,
	iterationNumber int,
	comp Comparator) {
	var averageTime float64

	for i := 0; i < iterationNumber; i++ {
		arr := generateRandomArray(size, limit)

		start := time.Now()
		sort(arr, comp)
		end := time.Since(start).Seconds()

		averageTime += end

		fmt.Printf("Sort with %d elements: %f seconds\n", size, end)
	}

	fmt.Printf("Average time: %f seconds\n", averageTime/float64(iterationNumber))
}
