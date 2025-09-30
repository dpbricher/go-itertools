package itertools

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ReduceSuite struct {
	suite.Suite
}

func (suite *ReduceSuite) TestReduce() {
	suite.Run("Reduce to string", func() {
		input := slices.Values([]int{1, 2, 3})
		result := Reduce(input, func(total string, item int) string {
			return fmt.Sprintf("%v%v", total, item)
		}, "0")

		suite.Equal("0123", result)
	})

	suite.Run("Reduce to slice", func() {
		input := slices.Values([]int{1, 2, 3})
		result := Reduce(input, func(total []int, item int) []int {
			return append(total, item)
		}, make([]int, 0, 3))

		suite.Equal([]int{1, 2, 3}, result)
	})

	suite.Run("Reduce to map", func() {
		input := slices.Values([]int{1, 2, 3})
		result := Reduce(input, func(total map[int]int, item int) map[int]int {
			total[item] = item
			return total
		}, make(map[int]int))

		suite.Equal(map[int]int{1: 1, 2: 2, 3: 3}, result)
	})
}

func (suite *ReduceSuite) TestReduceBreak() {
	input := slices.Values([]int{1, 2})
	suite.NotPanics(func() {
		for range Reduce(input, func(total, item int) int {
			return 0
		}, 0) {
			break
		}
	})
}

func (suite *ReduceSuite) TestReduce2() {
	suite.Run("Reduce to string", func() {
		input := slices.All([]int{1, 2, 3})
		result := Reduce2(input, func(total string, index, item int) string {
			return fmt.Sprintf("%s %d:%d", total, index, item)
		}, "#")

		suite.Equal("# 0:1 1:2 2:3", result)
	})

	suite.Run("Reduce to slice", func() {
		input := slices.All([]int{1, 2, 3})
		result := Reduce2(input, func(total []int, index, item int) []int {
			return append(total, index+item)
		}, make([]int, 0, 3))

		suite.Equal([]int{1, 3, 5}, result)
	})

	suite.Run("Reduce to map", func() {
		input := slices.All([]int{1, 2, 3})
		result := Reduce2(input, func(total map[int]int, index, item int) map[int]int {
			total[index] = item
			return total
		}, make(map[int]int))

		suite.Equal(map[int]int{0: 1, 1: 2, 2: 3}, result)
	})
}

func (suite *ReduceSuite) TestReduce2Break() {
	input := slices.All([]int{1, 2})
	suite.NotPanics(func() {
		for range Reduce2(input, func(total, index, item int) int {
			return 0
		}, 0) {
			break
		}
	})
}

func TestReduce(t *testing.T) {
	suite.Run(t, new(ReduceSuite))
}
