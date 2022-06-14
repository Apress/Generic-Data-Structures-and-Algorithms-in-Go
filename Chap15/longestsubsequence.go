package main 

import (
	"fmt"
)

func max(value1, value2 int) int {
	if value1 >= value2 {
		return value1
	} else {
		return value2
	}
}

func reverse(x []rune) []rune {
	result := []rune{}
	for index := len(x) - 1; index >= 0; index-- {
		result = append(result, x[index])
	}
	return result
}

func longestCommonSubsequenceTable(x, y []rune) (LCS [][]int) {
	// Return matrix so that LCS[j][k] is longest common
	// sequence for x[0:j] and y[0:k]
	n := len(x)
	m := len(y)
	// Initialize LCS table of size (n + 1 x m + 1)
	LCS = make([][]int, n + 1)
	for row := 0; row < n + 1; row++ {
		LCS[row] = make([]int, m + 1)
	}
	for row := 0; row < n; row++ {
		for col := 0; col < m; col++ {
			if x[row] == y[col] {
				LCS[row + 1][col + 1] = 1 + LCS[row][col]
			} else {
				LCS[row + 1][col + 1] = max(LCS[row][col + 1], LCS[row + 1][col])
			}
		}
	}
	return LCS
}

func LongestCommonSequence(x, y []rune) string {
	table := longestCommonSubsequenceTable(x, y)
	result := []rune{}
	j, k := len(x), len(y)
	for {
		if table[j][k] == 0 {
			break
		}
		if x[j - 1] == y[k - 1] {
			result = append(result, x[j - 1])
			j -= 1 
			k -= 1
		} else if table[j - 1][k] >= table[j][k - 1] {
			j -= 1 
			k -= 1

		}
	}
	return string(reverse(result))
}



func main() {
	x := "CGTTACAATTTGCG"
	y := "TTTTAAACGTGCG"
	lcs := LongestCommonSequence([]rune(x), []rune(y))
	fmt.Println(lcs)

	x = "ATCGAATTCCGGTAGTCGT"
	y = "CGATAGTTCAGCCAG"
	lcs = LongestCommonSequence([]rune(x), []rune(y))
	fmt.Println(lcs)

}
/* Output
TTAATGCG
TAGC
*/
