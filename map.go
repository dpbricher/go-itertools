package itertools

import (
	"iter"
)

// Map creates a new [iter.Seq] by applying the given mapper func to the given [iter.Seq].
//
//	input := slices.Values([]int{1, 2, 3})
//	output := itertools.Map(input, func(item int) string {
//	  return fmt.Sprintf("%d", item)
//	})
//
//	slices.Collect(output) // []string{"1", "2", "3"}
func Map[In, Out any](seq iter.Seq[In], mapper func(In) Out) iter.Seq[Out] {
	return func(yield func(Out) bool) {
		for item := range seq {
			if !yield(mapper(item)) {
				return
			}
		}
	}
}

// Map2 creates a new [iter.Seq2] by applying the given mapper func to the given [iter.Seq2].
//
//	 input := slices.All([]int{1, 2, 3})
//		output := itertools.Map2(input, func(index, item int) (int, string) {
//			return index, fmt.Sprintf("%d", item)
//		})
//
//		maps.Collect(output) // map[int]string{0: "1", 1: "2", 2: "3"}
func Map2[In1, In2, Out1, Out2 any](seq iter.Seq2[In1, In2], mapper func(In1, In2) (Out1, Out2)) iter.Seq2[Out1, Out2] {
	return func(yield func(Out1, Out2) bool) {
		for item1, item2 := range seq {
			if !yield(mapper(item1, item2)) {
				return
			}
		}
	}
}
