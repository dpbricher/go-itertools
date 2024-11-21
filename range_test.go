package itertools

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RangeSuite struct {
	suite.Suite
}

func (suite *RangeSuite) TestAscending() {
	suite.Run("+ve to +ve", func() {
		actualSlice := slices.Collect(Range(1, 3))
		suite.Equal([]int{1, 2, 3}, actualSlice)
	})

	suite.Run("-ve to +ve", func() {
		actualSlice := slices.Collect(Range(-1, 1))
		suite.Equal([]int{-1, 0, 1}, actualSlice)
	})

	suite.Run("-ve to -ve", func() {
		actualSlice := slices.Collect(Range(-3, -1))
		suite.Equal([]int{-3, -2, -1}, actualSlice)
	})
}

func (suite *RangeSuite) TestDescending() {
	suite.Run("+ve to +ve", func() {
		actualSlice := slices.Collect(Range(3, 1))
		suite.Equal([]int{3, 2, 1}, actualSlice)
	})

	suite.Run("+ve to -ve", func() {
		actualSlice := slices.Collect(Range(1, -1))
		suite.Equal([]int{1, 0, -1}, actualSlice)
	})

	suite.Run("-ve to -ve", func() {
		actualSlice := slices.Collect(Range(-1, -3))
		suite.Equal([]int{-1, -2, -3}, actualSlice)
	})
}

func (suite *RangeSuite) TestEqual() {
	suite.Run("+ve", func() {
		actualSlice := slices.Collect(Range(1, 1))
		suite.Equal([]int{1}, actualSlice)
	})

	suite.Run("Zero", func() {
		actualSlice := slices.Collect(Range(0, 0))
		suite.Equal([]int{0}, actualSlice)
	})

	suite.Run("-ve", func() {
		actualSlice := slices.Collect(Range(-1, -1))
		suite.Equal([]int{-1}, actualSlice)
	})
}

func TestRange(t *testing.T) {
	suite.Run(t, new(RangeSuite))
}

type RangeStepSuite struct {
	suite.Suite
}

func (suite *RangeStepSuite) TestAscending() {
	suite.Run("Including bounds", func() {
		actualSlice := slices.Collect(RangeStep(1, 5, 2))
		suite.Equal([]int{1, 3, 5}, actualSlice)
	})

	suite.Run("Excluding upper bound", func() {
		actualSlice := slices.Collect(RangeStep(1, 4, 2))
		suite.Equal([]int{1, 3}, actualSlice)
	})

	suite.Run("Step equals range", func() {
		actualSlice := slices.Collect(RangeStep(1, 5, 5))
		suite.Equal([]int{1}, actualSlice)
	})

	suite.Run("Start equals end", func() {
		actualSlice := slices.Collect(RangeStep(1, 1, 1))
		suite.Equal([]int{1}, actualSlice)
	})

	suite.Run("Zero step", func() {
		start := 1
		end := 5
		step := 0

		suite.PanicsWithValue(fmt.Sprintf(infiniteRangeMessageTemplate, start, step, end), func() {
			slices.Collect(RangeStep(start, end, step))
		})
	})

	suite.Run("Wrong step direction", func() {
		suite.Panics(func() {
			slices.Collect(RangeStep(1, 5, -1))
		})
	})
}

func (suite *RangeStepSuite) TestDescending() {
	suite.Run("Including bounds", func() {
		actualSlice := slices.Collect(RangeStep(5, 1, -2))
		suite.Equal([]int{5, 3, 1}, actualSlice)
	})

	suite.Run("Excluding upper bound", func() {
		actualSlice := slices.Collect(RangeStep(5, 2, -2))
		suite.Equal([]int{5, 3}, actualSlice)
	})

	suite.Run("Step equals range", func() {
		actualSlice := slices.Collect(RangeStep(5, 1, -5))
		suite.Equal([]int{5}, actualSlice)
	})

	suite.Run("Start equals end", func() {
		actualSlice := slices.Collect(RangeStep(1, 1, -1))
		suite.Equal([]int{1}, actualSlice)
	})

	suite.Run("Zero step", func() {
		suite.Panics(func() {
			slices.Collect(RangeStep(5, 1, 0))
		})
	})

	suite.Run("Wrong step direction", func() {
		suite.Panics(func() {
			slices.Collect(RangeStep(5, 1, 1))
		})
	})
}

func TestRangeStep(t *testing.T) {
	suite.Run(t, new(RangeStepSuite))
}
