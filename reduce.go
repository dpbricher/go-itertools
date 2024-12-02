package itertools

import (
	"iter"
)

// Reduce applies the given collector func to every element of the given [iter.Seq] and returns the result.
//
// For the first item in `seq`, the `collector` func will be called with two arguments; the value of `initial` and the item from `seq`. For every item in `seq` after the first, the `collector` func will be called with the output from the previous call to `collector` and the item from `seq`.
//
// The intent behind `collector` is that it takes each item from the sequence and adds them to an accumulated total; the first argument to `collector` is the total so far, and the value returned from `collector` is the new total after the current item has been added into it. `initial` is used to give this total a starting value.
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
// This func works the same as [Reduce], except that the input iterator is [iter.Seq2].
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
