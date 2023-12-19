package main

func finalValueAfterOperations(operations []string) int {
	result := 0
    for i := 0; i < len(operations); i++ {
		for j := 0; j < len(operations[i]); j++ {
			if operations[i][j] == '+' {
				result++
				continue
			} else if operations[i][j] == '-' {
				result--
				continue
			}
		}
	}
	return result
}