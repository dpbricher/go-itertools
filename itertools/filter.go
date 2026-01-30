package itertools

import (
	"iter"
)

// Filter wraps an [iter.Seq] and drops items that do not match a supplied filtering func.
//
// For every item in seq, filter will be called and passed that item as its sole argument. If that call to filter returns true then that item will be included in the output [iter.Seq], otherwise it will be discarded.
//
// The order of items in the output [iter.Seq] will match the order that they appear in seq.
//
//	input := slices.Values([]int{0, 1, 2, 3})
//	filteredIter := itertools.Filter(input, func(item int) bool {
//		return item%2 == 0
//	})
//
//	slices.Collect(filteredIter) // []int{0, 2}
func Filter[T any](seq iter.Seq[T], filter func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range seq {
			if filter(item) {
				if !yield(item) {
					return
				}
			}
		}
	}
}

// Filter2 wraps an [iter.Seq2] and drops items that do not match a supplied filtering func.
//
// This func works the same as [Filter], except that the input and output iterators are [iter.Seq2].
//
//	input := slices.All([]int{0, 1, 2, 3})
//	filteredIter := itertools.Filter2(input, func(index, item int) bool {
//		return index == 0 || item == 3
//	})
//
//	maps.Collect(filteredIter) // map[int]int{0: 0, 3: 3}
func Filter2[S, T any](seq iter.Seq2[S, T], filter func(S, T) bool) iter.Seq2[S, T] {
	return func(yield func(S, T) bool) {
		for item1, item2 := range seq {
			if filter(item1, item2) {
				if !yield(item1, item2) {
					return
				}
			}
		}
	}
}
