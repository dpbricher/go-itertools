package itertools

import (
	"fmt"
	"iter"
)

var infiniteRangeMessageTemplate = "iterating over an infinite range (start: %d, step: %d, end: %d)"

// Range creates an [iter.Seq] that will return all integers between `start` and `end`, including `start` and `end`.
//
//	slices.Collect(itertools.Range(1, 5)) // []int{1, 2, 3, 4, 5}
func Range(start, end int) iter.Seq[int] {
	if start < end {
		return rangeAscending(start, end, 1)
	}

	return rangeDescending(start, end, -1)
}

// RangeStep creates an [iter.Seq] that will return all integers between `start` and `end`, in increments of `step`.
//
//	slices.Collect(itertools.Range(1, 5, 2)) // []int{1, 3, 5}
//
// `start` will always be included as the first value, but `end` will only be included if the distance between `start` and `end` is exactly divisable by `step`.
//
//	slices.Collect(itertools.Range(5, 2, -2)) // []int{5, 3}
//
// If the specified `step` would result in an infinite iterator then the returned iterator will panic the first time that a value is read from it.
//
//	slices.Collect(itertools.Range(1, 5, -1)) // panics
func RangeStep(start, end, step int) iter.Seq[int] {
	if start < end {
		return rangeAscending(start, end, step)
	}

	if start == end {
		return func(yield func(int) bool) {
			yield(start)
		}
	}

	return rangeDescending(start, end, step)
}

func infiniteRangeMessage(start, end, step int) string {
	return fmt.Sprintf(infiniteRangeMessageTemplate, start, step, end)
}

func rangeAscending(start, end, step int) iter.Seq[int] {
	return func(yield func(int) bool) {
		// if `step` is not a positive value then this method would loop infinitely
		if step < 1 {
			panic(infiniteRangeMessage(start, end, step))
		}

		for i := start; i <= end; i += step {
			if !yield(i) {
				return
			}
		}
	}
}

func rangeDescending(start, end, step int) iter.Seq[int] {
	return func(yield func(int) bool) {
		// if `step` is not a negative value then this method would loop infinitely
		if step > -1 {
			panic(infiniteRangeMessage(start, end, step))
		}

		for i := start; i >= end; i += step {
			if !yield(i) {
				return
			}
		}
	}
}
