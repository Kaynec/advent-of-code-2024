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

func appendIfInBound(cord []int, locs *[][]int, slice [][]string, table map[string]bool) {
	if cord[1] == 0 {
		fmt.Println("ZERO BASED")
	}
	tableTemplate := fmt.Sprintf("%d:%d", cord[0], cord[1])
	_, alreadyExist := table[tableTemplate]
	if cord[0] < len(slice) && cord[0] >= 0 && cord[1] < len(slice[0]) && cord[1] >= 0 && !alreadyExist {
		*locs = append(*locs, cord)
		table[tableTemplate] = true
	}
}

func walk(cord []int, nextCord []int, table map[string]bool, strSlice [][]string, goodLocs [][]int) {
	rowDiff := cord[0] - nextCord[0]
	colDiff := cord[1] - nextCord[1]

	for k := 0; k < 1000000000000000000; k++ {
		cord := []int{rowDiff + cord[0], colDiff + cord[1]}
		if cord[0] > len(strSlice) || cord[0] < 0 || cord[1] > len(strSlice[0]) || cord[1] < 1 {
			break
		}
		appendIfInBound(cord, &goodLocs, strSlice, table)

		rowDiff = (cord[0] - nextCord[0]) * k
		colDiff = (cord[1] - nextCord[1]) * k
	}
}
func getGoodLocsCount(strSlice [][]string, char string, table map[string]bool) [][]int {
	charLocs := [][]int{}
	goodLocs := [][]int{}
	for lineIndex, line := range strSlice {
		for el := range line {

			if strSlice[lineIndex][el] == char {
				charLocs = append(charLocs, []int{lineIndex, el})
				appendIfInBound([]int{lineIndex, el}, &goodLocs, strSlice, table)
			}
		}
	}
	for i := 0; i < len(charLocs)-1; i++ {

		for j := i; j < len(charLocs)-1; j++ {
			curr := charLocs[i]
			next := charLocs[j+1]

			walk(curr, next, table, strSlice, goodLocs)
			walk(next, curr, table, strSlice, goodLocs)

		}

	}

	return goodLocs
}

func getAllChars(slice [][]string) (res []string) {
	table := make(map[string]bool)

	for row := range slice {
		for col := range slice[row] {
			if slice[row][col] != "." && !table[slice[row][col]] {

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

	// SORTING MIGHT BE NECCACARY
	// sort.Slice(goodLocs, func(i, j int) bool { return goodLocs[i][0] < goodLocs[j][0] })

	fmt.Print(goodLocs, len(goodLocs))

}
func main() {

	path := "sample"
	if os.Args[1:] != nil && os.Args[1:][0] != "" {
		path = os.Args[1:][0]
	}
	partOne(path)
}
