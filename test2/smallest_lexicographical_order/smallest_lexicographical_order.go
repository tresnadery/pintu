package smallest_lexicographical_order

import (
	"strings"
)

func SmallestLexicographicalOrder(s string) string {
	chars := strings.Split(s, "")
	size := len(s)
	for i := 1; i < size; i++ {
		j := i
		for j > 0 {
			if chars[j-1] > chars[j] {
				chars[j-1], chars[j] = string(chars[j]), string(chars[j-1])
			}
			j = j - 1
		}
	}
	return strings.Join(chars, "")
}
