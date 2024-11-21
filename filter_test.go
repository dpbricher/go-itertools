package itertools

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/suite"
)

type FilterSuite struct {
	suite.Suite
}

func (suite *FilterSuite) TestFilter() {
	input := slices.Values([]int{0, 1, 2, 3})
	actualSlice := slices.Collect(
		Filter(input, func(item int) bool {
			return item%2 == 0
		}),
	)

	suite.Equal([]int{0, 2}, actualSlice)
}

func (suite *FilterSuite) TestFilterMultiUse() {
	input := slices.Values([]int{0, 1, 2, 3})
	iterator := Filter(input, func(item int) bool {
		return item%2 != 0
	})

	firstSlice := slices.Collect(iterator)
	secondSlice := slices.Collect(iterator)

	suite.Equal([]int{1, 3}, firstSlice)
	suite.Equal(firstSlice, secondSlice)
}

func (suite *FilterSuite) TestFilter2() {
	input := slices.All([]int{0, 1, 2, 3})
	actualSlice := maps.Collect(
		Filter2(input, func(index, item int) bool {
			return index == 0 || item == 3
		}),
	)

	suite.Equal(map[int]int{0: 0, 3: 3}, actualSlice)
}

func (suite *FilterSuite) TestFilter2MultiUse() {
	input := slices.All([]int{0, 1, 2, 3})
	iterator := Filter2(input, func(index, item int) bool {
		return index == 1 || item == 2
	})

	firstSlice := maps.Collect(iterator)
	secondSlice := maps.Collect(iterator)

	suite.Equal(map[int]int{1: 1, 2: 2}, firstSlice)
	suite.Equal(firstSlice, secondSlice)
}

func TestFilter(t *testing.T) {
	suite.Run(t, new(FilterSuite))
}
