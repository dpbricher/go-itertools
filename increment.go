package itertools

import "iter"

// Asc creates an [iter.Seq] that will output an ascending sequence of integers.
//
// The first item in the sequence will be `start`, then each item after will be `step` more than the previous one. The iterator will output a total of `count` values.
//
// If specifying the end value as an argument is desired (instead of the number of outputs), consider using [RangeStep] instead.
//
//	input := Asc(1, 1, 3)
//	output := slices.Collect(input) // []int{1, 2, 3}
func Asc(start int, step, count uint) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := uint(0); i < count; i += step {
			if !yield(start + int(i)) {
				return
			}
		}
	}
}

// Desc creates an [iter.Seq] that will output a descending sequence of integers.
//
// The first item in the sequence will be `start`, then each item after will be `step` less than the previous one. The iterator will output a total of `count` values.
//
// If specifying the end value as an argument is desired (instead of the number of outputs), consider using [RangeStep] instead.
//
//	input := Desc(3, 1, 3)
//	output := slices.Collect(input) // []int{3, 2, 1}
func Desc(start int, step, count uint) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := uint(0); i < count; i += step {
			if !yield(start - int(i)) {
				return
			}
		}
	}
}
