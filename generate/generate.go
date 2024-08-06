package generate

func GenerateCombinations(prefix string, n int, result *[]string) {
	if len(prefix) == n {
		*result = append(*result, prefix)
		return
	}
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for _, char := range chars {
		newPrefix := prefix + string(char)
		GenerateCombinations(newPrefix, n, result)
	}
}
