package profit_calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateProfit_WithEmptyParameter(t *testing.T) {
	profit, err := CalculateProfit([]string{})
	if assert.Error(t, err) {
		assert.Equal(t, "prices is empty", err.Error())
	}
	assert.Equal(t, 0, profit)
}

func TestCalculateProfit_WithZeroProfit(t *testing.T) {
	prices := []string{"5", "4", "3", "2", "1"}
	profit, err := CalculateProfit(prices)
	assert.Equal(t, 0, profit)
	assert.Nil(t, err)
}

func TestCalculateProfit_WithProfit(t *testing.T) {
	prices := []string{"3", "2", "1", "5", "6", "2"}
	profit, err := CalculateProfit(prices)
	assert.Equal(t, 5, profit)
	assert.Nil(t, err)
}
