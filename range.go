package itertools

import (
	"iter"
)

// Range creates an [iter.Seq] that will output all integers between start and end, including start and end.
//
//	slices.Collect(itertools.Range(1, 5)) // []int{1, 2, 3, 4, 5}
func Range(start, end int) iter.Seq[int] {
	if start < end {
		return rangeAscending(start, end, 1)
	}

	return rangeDescending(start, end, -1)
}

// RangeStep creates an [iter.Seq] that will output all integers between start and end, in increments of step.
//
//	slices.Collect(itertools.Range(1, 5, 2)) // []int{1, 3, 5}
//
// start will always be included as the first value, but end will only be included if the distance between start and end is exactly divisable by step.
//
//	slices.Collect(itertools.Range(5, 2, -2)) // []int{5, 3}
//
// If step is not a value that moves start closer to end, then the returned iterator will be empty. This is to prevent creation of an iterator that is incrementing in the wrong direction, or that is infinite.
//
//	slices.Collect(itertools.Range(1, 5, -1)) // []int(nil)
//	slices.Collect(itertools.Range(1, 5, 0)) // []int(nil)
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

func emptyIterator(_ func(int) bool) {
}

func rangeAscending(start, end, step int) iter.Seq[int] {
	// if step is not a positive value then this method is probably looping the wrong way
	if step < 1 {
		return emptyIterator
	}

	return func(yield func(int) bool) {
		for i := start; i <= end; i += step {
			if !yield(i) {
				return
			}
		}
	}
}

func rangeDescending(start, end, step int) iter.Seq[int] {
	// if step is not a negative value then this method is probably looping the wrong way
	if step > -1 {
		return emptyIterator
	}

	return func(yield func(int) bool) {
		for i := start; i >= end; i += step {
			if !yield(i) {
				return
			}
		}
	}
}
