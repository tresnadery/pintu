package calculate_number_of_factor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateNumberOfFactor_With128AsParameter(t *testing.T) {
	total := GetCountOfNumberWithSixFactor(128)
	assert.Equal(t, total, 19)
}

func TestCalculateNumberOfFactor_With1024AsParameter(t *testing.T) {
	total := GetCountOfNumberWithSixFactor(1024)
	assert.Equal(t, total, 112)
}

func TestCalculateNumberOfFactor_With16384AsParameter(t *testing.T) {
	total := GetCountOfNumberWithSixFactor(16384)
	assert.Equal(t, total, 1168)
}
