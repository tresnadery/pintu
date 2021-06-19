package main

import (
	"fmt"
	"pintu/test3/calculate_number_of_factor"
)

func main() {
	n := 262144
	total := calculate_number_of_factor.GetCountOfNumberWithSixFactor(n)
	fmt.Printf("H(262144) = %d", total)
}
