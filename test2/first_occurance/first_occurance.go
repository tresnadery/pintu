package first_occurance

func FirstOccurance(s string) string {
	res := ""
	arrChar := map[string]int{}
	for _, char := range s {
		charConv := string(char)
		arrChar[charConv]++
		if arrChar[charConv] == 1 {
			res += charConv
		}
	}
	return res
}
