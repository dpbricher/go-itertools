package itertools

import (
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MapSuite struct {
	suite.Suite
}

func (suite *MapSuite) TestMap() {
	input := slices.Values([]int{1, 2, 3})
	output := slices.Collect(
		Map(input, func(item int) string {
			return fmt.Sprintf("%d", item)
		}),
	)

	suite.Equal([]string{"1", "2", "3"}, output)
}

func (suite *MapSuite) TestMap2() {
	input := slices.All([]int{1, 2, 3})
	output := maps.Collect(
		Map2(input, func(index, item int) (int, string) {
			return index, fmt.Sprintf("%d", item)
		}),
	)

	suite.Equal(map[int]string{0: "1", 1: "2", 2: "3"}, output)
}

func TestMap(t *testing.T) {
	suite.Run(t, new(MapSuite))
}
