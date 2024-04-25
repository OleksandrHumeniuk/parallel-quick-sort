package main

import "fmt"

const (
	ARRAY_LIMIT = 100
	ARRAY_SIZE  = 1_000_000

	TEST_SEQUENTIAL = false
	TIME_SEQUENTIAL = false

	TEST_PARALLEL      = false
	TIME_PARALLEL      = false
	PARAMETER_PARALLEL = true
	MAX_PARALLEL_DEPTH = 50

	RUNS_PER_FUNCTION = 10
)

func main() {

	if TEST_SEQUENTIAL {
		fmt.Println("Sequential Quick Sort")
		testAlgorithm(sequentialQuickSortStart, ARRAY_SIZE, ARRAY_LIMIT, RUNS_PER_FUNCTION)
	}

	if TIME_SEQUENTIAL {
		for i := 0; i < RUNS_PER_FUNCTION; i++ {
			fmt.Println("\n\nSequential Quick Sort")

			size := ARRAY_SIZE * (i + 1)

			timeAlgorithm(sequentialQuickSortStart, size, ARRAY_LIMIT, RUNS_PER_FUNCTION)
		}
	}

	if TEST_PARALLEL {
		fmt.Println("Parallel Quick Sort")
		testAlgorithm(func(arr []int) []int {
			return parallelQuickSortStart(arr, MAX_PARALLEL_DEPTH)
		}, ARRAY_SIZE, ARRAY_LIMIT, RUNS_PER_FUNCTION)
	}

	if TIME_PARALLEL {
		for i := 0; i < RUNS_PER_FUNCTION; i++ {
			fmt.Println("\n\nParallel Quick Sort")

			size := ARRAY_SIZE * (i + 1)

			timeAlgorithm(func(arr []int) []int {
				return parallelQuickSortStart(arr, MAX_PARALLEL_DEPTH)
			}, size, ARRAY_LIMIT, RUNS_PER_FUNCTION)
		}
	}

	if PARAMETER_PARALLEL {
		for i := 1; i < 100; i++ {
			fmt.Println("\n\nParallel Quick Sort with maximum depth", i)

			timeAlgorithm(func(arr []int) []int {
				return parallelQuickSortStart(arr, i)
			}, ARRAY_SIZE, ARRAY_LIMIT, RUNS_PER_FUNCTION)
		}
	}
}
