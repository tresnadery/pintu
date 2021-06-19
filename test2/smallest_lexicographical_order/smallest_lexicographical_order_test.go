package smallest_lexicographical_order

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestLexicographicalOrder_WithEmptyString(t *testing.T) {
	res := SmallestLexicographicalOrder("")
	assert.Equal(t, res, "")
}
func TestSmallestLexicographicalOrderTest_WithRightString(t *testing.T) {
	res := SmallestLexicographicalOrder("sebar")
	assert.Equal(t, res, "abers")
}

func TestSmallestLexicographicalOrder_WithSecondRightString(t *testing.T) {
	res := SmallestLexicographicalOrder("charlie")
	assert.Equal(t, res, "acehilr")
}
