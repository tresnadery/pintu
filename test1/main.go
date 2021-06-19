package main

import (
	"fmt"
	"log"
	"pintu/test1/profit_calculation"
)

func main() {
	historyPrice, err := profit_calculation.GetHistoricalPrices()
	if err != nil {
		log.Fatal(err)
	}
	profit, err := profit_calculation.CalculateProfit(historyPrice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the maximum profit Jacky can make is: $%d", profit)
}
