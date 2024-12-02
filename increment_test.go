package itertools

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/suite"
)

type IncrementSuite struct {
	suite.Suite
}

// TODO: add tests when step > 1
func (suite *IncrementSuite) TestAsc() {
	suite.Run("+ve to +ve", func() {
		input := Asc(1, 1, 3)
		output := slices.Collect(input)

		suite.Equal([]int{1, 2, 3}, output)
	})

	suite.Run("-ve to +ve", func() {
		input := Asc(-1, 1, 3)
		output := slices.Collect(input)

		suite.Equal([]int{-1, 0, 1}, output)
	})

	suite.Run("-ve to -ve", func() {
		input := Asc(-3, 1, 3)
		output := slices.Collect(input)

		suite.Equal([]int{-3, -2, -1}, output)
	})
}

// TODO: add tests when step > 1
func (suite *IncrementSuite) TestDesc() {
	suite.Run("+ve to +ve", func() {
		input := Desc(3, 1, 3)
		output := slices.Collect(input)

		suite.Equal([]int{3, 2, 1}, output)
	})

	suite.Run("+ve to -ve", func() {
		input := Desc(1, 1, 3)
		output := slices.Collect(input)

		suite.Equal([]int{1, 0, -1}, output)
	})

	suite.Run("-ve to -ve", func() {
		input := Desc(-1, 1, 3)
		output := slices.Collect(input)

		suite.Equal([]int{-1, -2, -3}, output)
	})
}

func TestIncrement(t *testing.T) {
	suite.Run(t, new(IncrementSuite))
}
