package main

import "sync"

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func sequentialQuickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = sequentialQuickSort(arr, low, p-1)
		arr = sequentialQuickSort(arr, p+1, high)
	}
	return arr
}

func parallelQuickSort(arr []int, low, high int, wg *sync.WaitGroup, depth int) {
	defer wg.Done()

	if low < high {
		var p int
		arr, p = partition(arr, low, high)

		if depth > 0 {
			var innerWg sync.WaitGroup
			innerWg.Add(2)
			go parallelQuickSort(arr, low, p-1, &innerWg, depth-1)
			go parallelQuickSort(arr, p+1, high, &innerWg, depth-1)
			innerWg.Wait()
		} else {
			sequentialQuickSort(arr, low, p-1)
			sequentialQuickSort(arr, p+1, high)
		}
	}
}

func sequentialQuickSortStart(arr []int) []int {
	return sequentialQuickSort(arr, 0, len(arr)-1)
}

func parallelQuickSortStart(arr []int, maxDepth int) []int {
	var wg sync.WaitGroup
	wg.Add(1)
	go parallelQuickSort(arr, 0, len(arr)-1, &wg, maxDepth)
	wg.Wait()
	return arr
}
