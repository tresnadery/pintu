package calculate_number_of_factor

func GetCountOfNumberWithSixFactor(n int) int {
	var numberWithSicFactor int
	for i := 1; i <= n; i++ {
		isOnlySixFactor := countFactor(i)
		if isOnlySixFactor {
			numberWithSicFactor++
		}
	}
	return numberWithSicFactor
}

func countFactor(n int) bool {
	arr := map[int]int{}
	j := n
	i := 1
	for {
		if n%i != 0 {
			i++
			continue
		}
		if i >= j {
			break
		}
		if n%i == 0 {
			j = n / i
			arr[i]++
			arr[j]++
		}
		i++
	}
	return len(arr) == 6
}
