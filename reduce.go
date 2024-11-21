package itertools

import (
	"iter"
)

// Reduce applies the given collector func to every element of the given [iter.Seq] and returns the result.
//
//	 input := slices.Values([]int{1, 2, 3})
//		result := itertools.Reduce(input, func(total []int, item int) []int {
//			return append(total, item)
//		}, make([]int, 0, 3))
//
//		// result = []int{1, 2, 3}
func Reduce[In, Out any](seq iter.Seq[In], collector func(Out, In) Out, initial Out) Out {
	reduced := initial
	for item := range seq {
		reduced = collector(reduced, item)
	}
	return reduced
}

// Reduce2 applies the given collector func to every element of the given [iter.Seq2] and returns the result.
//
//	 input := slices.All([]int{1, 2, 3})
//		result := itertools.Reduce2(input, func(total []int, index, item int) []int {
//			return append(total, index+item)
//		}, make([]int, 0, 3))
//
//		// result = []int{1, 3, 5}
func Reduce2[In1, In2, Out any](seq iter.Seq2[In1, In2], collector func(Out, In1, In2) Out, initial Out) Out {
	reduced := initial
	for item1, item2 := range seq {
		reduced = collector(reduced, item1, item2)
	}
	return reduced
}
