package timesort

import (
	"math"
)

// Comparer interface is use to compare two elements.
// It has only one method Copmare which return int value: greater than 0 if element passed as object is greater, 0 if elements are equal, and negative value if element is greater.
type Comparer interface {
	Compare(b interface{}) int
}

func binarySearch(data []Comparer, item Comparer, start int, end int) int {
	if start == end {
		if data[start].Compare(item) > 0 {
			return start
		}
		return start + 1
	}

	if start > end {
		return start
	}

	mid := int(math.Round((float64(start + end)) / 2))
	if data[mid].Compare(item) < 0 {
		return binarySearch(data, item, mid+1, end)
	}
	if data[mid].Compare(item) > 0 {
		return binarySearch(data, item, start, mid-1)
	}

	return mid
}

func insertionSort(data []Comparer) []Comparer {
	for i := 1; i < len(data); i++ {
		var result []Comparer
		value := data[i]
		pos := binarySearch(data, value, 0, i-1)

		result = append(result, data[0:pos]...)
		result = append(result, value)
		result = append(result, data[pos:i]...)
		result = append(result, data[i+1:]...)

		data = result
	}

	return data
}

func merge(left []Comparer, right []Comparer) []Comparer {
	var result []Comparer

	if len(left) == 0 {
		return right
	}
	if len(right) == 0 {
		return left
	}

	if left[0].Compare(right[0]) < 0 {
		result = merge(left[1:], right)
		result = append([]Comparer{left[0]}, result...)
	} else {
		result = merge(left, right[1:])
		result = append([]Comparer{right[0]}, result...)
	}

	return result
}

// Timesort sort elements (implementing Comparer interface) using Timesort algorithm.
func Timesort(elements []Comparer) []Comparer {
	if len(elements) < 2 {
		return elements
	}

	var runs [][]Comparer
	var sortedRuns [][]Comparer
	newRun := []Comparer{elements[0]}

	for i := 1; i < len(elements); i++ {
		if i == len(elements)-1 {
			newRun = append(newRun, elements[i])
			runs = append(runs, newRun)
			break
		}

		if elements[i].Compare(elements[i-1]) < 0 {
			runs = append(runs, newRun)
			newRun = []Comparer{elements[i]}
		} else {
			newRun = append(newRun, elements[i])
		}
	}

	for _, item := range runs {
		sortedRuns = append(sortedRuns, insertionSort(item))
	}
	var sortedArray []Comparer
	for _, item := range sortedRuns {
		sortedArray = merge(sortedArray, item)
	}

	return sortedArray
}
