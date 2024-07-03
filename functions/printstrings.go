package functions

import "fmt"

func ProcessString(slice []int, filename string) []string {
	var results []string
	var str string
	for _, number := range slice {
		str = Maps(number, filename)
		results = append(results, str)
	}

	return results
}

// slice the string into chunks of eight depending on the number of letters passed.And print them

func PrintStrings(str []string) {
	chunksize := 8
	var result [][]string
	for i := 0; i < len(str); i += chunksize {
		end := i + chunksize
		if end > len(str) {
			end = len(str)
		}
		result = append(result, str[i:end])
	}
	for i := 0; i < len(result); {
		for j := 0; j < 8; {
			fmt.Print(result[i][j])
			i++
			if i == len(result) {
				fmt.Println()
				j++
				i = 0
			}
		}
		break
	}

}
