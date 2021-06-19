package first_occurance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstOccurance_WithEmptyString(t *testing.T) {
	res := FirstOccurance("")
	assert.Equal(t, res, "")
}

func TestFirstOccurance_WithRightString(t *testing.T) {
	res := FirstOccurance("sebaerb")
	assert.Equal(t, res, "sebar")
}
