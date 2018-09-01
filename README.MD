[![Go Report Card](https://goreportcard.com/badge/github.com/mimol91/timesort)](https://goreportcard.com/report/github.com/mimol91/timesort)
[![Build Status](https://travis-ci.org/mimol91/timesort.svg?branch=master)](https://travis-ci.org/mimol91/timesort)
[![codecov](https://codecov.io/gh/mimol91/timesort/branch/master/graph/badge.svg)](https://codecov.io/gh/mimol91/timesort)
# Timesort

Timsort is a hybrid stable sorting algorithm, derived from merge sort and insertion sort, designed to perform well on many kinds of real-world data.

To use it create type implementing Comparer interface like:
```go
type comparableInt int

func (i comparableInt) Compare(b interface{}) int {
	return int(i - b.(comparableInt))
}
```

and invoke Timesort method.
```go
sortedElements := Timesort([]Comparer{comparableInt(2), comparableInt(-2)})
```
