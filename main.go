package main

import "fmt"

const (
	ARRAY_LIMIT = 100
	ARRAY_SIZE  = 1000

	TEST_SEQUENTIAL = false
	TIME_SEQUENTIAL = false

	TEST_PARALLEL      = true
	TIME_PARALLEL      = false
	PARAMETER_PARALLEL = false
	MAX_PARALLEL_DEPTH = 50

	RUNS_PER_FUNCTION = 10
)

func intComparator(a, b interface{}) bool {
	return a.(int) < b.(int)
}

func main() {

	if TEST_SEQUENTIAL {
		fmt.Println("Sequential Quick Sort")
		testAlgorithm(sequentialQuickSortStart, ARRAY_SIZE, ARRAY_LIMIT, RUNS_PER_FUNCTION, intComparator)
	}

	if TIME_SEQUENTIAL {
		for i := 0; i < RUNS_PER_FUNCTION; i++ {
			fmt.Println("\n\nSequential Quick Sort")

			size := ARRAY_SIZE * (i + 1)

			timeAlgorithm(sequentialQuickSortStart, size, ARRAY_LIMIT, RUNS_PER_FUNCTION, intComparator)
		}
	}

	if TEST_PARALLEL {
		fmt.Println("Parallel Quick Sort")
		testAlgorithm(func(arr []interface{}, comp Comparator) []interface{} {
			return parallelQuickSortStart(arr, MAX_PARALLEL_DEPTH, comp)
		}, ARRAY_SIZE, ARRAY_LIMIT, RUNS_PER_FUNCTION, intComparator)
	}

	if TIME_PARALLEL {
		for i := 0; i < RUNS_PER_FUNCTION; i++ {
			fmt.Println("\n\nParallel Quick Sort")

			size := ARRAY_SIZE * (i + 1)

			timeAlgorithm(func(arr []interface{}, comp Comparator) []interface{} {
				return parallelQuickSortStart(arr, MAX_PARALLEL_DEPTH, comp)
			}, size, ARRAY_LIMIT, RUNS_PER_FUNCTION, intComparator)
		}
	}

	if PARAMETER_PARALLEL {
		for i := 1; i < 100; i++ {
			fmt.Println("\n\nParallel Quick Sort with maximum depth", i)

			timeAlgorithm(func(arr []interface{}, comp Comparator) []interface{} {
				return parallelQuickSortStart(arr, i, comp)
			}, ARRAY_SIZE, ARRAY_LIMIT, RUNS_PER_FUNCTION, intComparator)
		}
	}
}
