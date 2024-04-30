package main

import "sync"

type Comparator func(a, b interface{}) bool

func partition(arr []interface{}, low, high int, comp Comparator) ([]interface{}, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if comp(arr[j], pivot) {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func sequentialQuickSort(arr []interface{}, low, high int, comp Comparator) []interface{} {
	if low < high {
		var p int
		arr, p = partition(arr, low, high, comp)
		arr = sequentialQuickSort(arr, low, p-1, comp)
		arr = sequentialQuickSort(arr, p+1, high, comp)
	}
	return arr
}

func parallelQuickSort(arr []interface{}, low, high int, wg *sync.WaitGroup, depth int, comp Comparator) {
	defer wg.Done()

	if low < high {
		var p int
		arr, p = partition(arr, low, high, comp)

		if depth > 0 {
			var innerWg sync.WaitGroup
			innerWg.Add(2)
			go parallelQuickSort(arr, low, p-1, &innerWg, depth-1, comp)
			go parallelQuickSort(arr, p+1, high, &innerWg, depth-1, comp)
			innerWg.Wait()
		} else {
			sequentialQuickSort(arr, low, p-1, comp)
			sequentialQuickSort(arr, p+1, high, comp)
		}
	}
}

func sequentialQuickSortStart(arr []interface{}, comp Comparator) []interface{} {
	return sequentialQuickSort(arr, 0, len(arr)-1, comp)
}

func parallelQuickSortStart(arr []interface{}, maxDepth int, comp Comparator) []interface{} {
	var wg sync.WaitGroup
	wg.Add(1)
	go parallelQuickSort(arr, 0, len(arr)-1, &wg, maxDepth, comp)
	wg.Wait()
	return arr
}
