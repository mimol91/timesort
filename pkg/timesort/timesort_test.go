package timesort

import (
	"testing"
)

type comparableInt int

func (i comparableInt) Compare(b interface{}) int {
	return int(i - b.(comparableInt))
}

func isEqual(elementsA []Comparer, elementsB []Comparer) bool {
	if len(elementsA) != len(elementsB) {
		return false
	}
	for i, value := range elementsA {
		if value != elementsB[i] {
			return false
		}
	}

	return true
}

func castIntToComparable(elements []int) []Comparer {
	var result []Comparer
	for _, element := range elements {
		result = append(result, comparableInt(element))
	}

	return result
}

func TestSort(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{2, 1, -2, 0, 0, 2, 72, -33}, []int{-33, -2, 0, 0, 1, 2, 2, 72}},
		{[]int{0}, []int{0}},
	}

	for _, test := range tests {
		input := castIntToComparable(test.input)
		expected := castIntToComparable(test.expected)
		actual := Timesort(input)

		if isEqual(actual, expected) == false {
			t.Errorf("Sort result for %d is not correct. Expected %d, has %d ", test.input, test.expected, actual)
		}
	}
}
