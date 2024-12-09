package main

import (
	"fmt"
	"os"
	"strings"
)

func parseStr(path string) [][]string {
	val, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := string(val)

	var sliced [][]string

	mulSlice := strings.Split(str, "\r\n")

	for _, rowVal := range mulSlice {

		sliced = append(sliced, strings.Split(rowVal, ""))
	}

	return sliced
}

func appendIfInBound(cord []int, cord2 []int, locs *[][]int, slice [][]string, table map[string]bool) {
	tableTemplate := fmt.Sprintf("%d:%d", cord[0], cord[1])
	_, alreadyExist := table[tableTemplate]
	if cord[0] < len(slice) && cord[0] >= 0 && cord[1] < len(slice[0]) && cord[1] >= 0 && !alreadyExist {
		*locs = append(*locs, cord)
		table[tableTemplate] = true
	}
}
func getGoodLocsCount(strSlice [][]string, char string, table map[string]bool) [][]int {
	charLocs := [][]int{}
	goodLocs := [][]int{}
	for lineIndex, line := range strSlice {
		for el := range line {

			if strSlice[lineIndex][el] == char {
				charLocs = append(charLocs, []int{lineIndex, el})
			}
		}
	}
	for i := 0; i < len(charLocs)-1; i++ {

		for j := i; j < len(charLocs)-1; j++ {
			curr := charLocs[i]
			next := charLocs[j+1]
			firstCor := []int{curr[0] - next[0] + curr[0], curr[1] - next[1] + curr[1]}
			secondCord := []int{next[0] - curr[0] + next[0], next[1] - curr[1] + next[1]}
			appendIfInBound(firstCor, secondCord, &goodLocs, strSlice, table)
			appendIfInBound(secondCord, firstCor, &goodLocs, strSlice, table)
		}

	}

	return goodLocs
}

func getAllChars(slice [][]string) (res []string) {
	table := make(map[string]bool)

	for row := range slice {
		for col := range slice[row] {
			if slice[row][col] == "." {
				continue
			}
			if !table[slice[row][col]] {
				res = append(res, slice[row][col])
			}
			table[slice[row][col]] = true
		}
	}

	return res
}
func partOne(path string) {
	table, strSlice, goodLocs := make(map[string]bool), parseStr(path), [][]int{}

	for _, char := range getAllChars(strSlice) {
		goodLocs = append(goodLocs, getGoodLocsCount(strSlice, char, table)...)
	}

	fmt.Print(len(goodLocs), goodLocs)

}
func main() {

	path := "sample"
	if os.Args[1:] != nil && os.Args[1:][0] != "" {
		path = os.Args[1:][0]
	}
	partOne(path)
}
