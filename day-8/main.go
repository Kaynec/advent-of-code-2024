package main

import (
	"fmt"
	"os"
	"strings"
)

const UNREACHABLE_NUMBER = 1000

func parseStr(path string) [][]string {
	val, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := string(val)

	var lineMatrix [][]string

	lines := strings.Split(str, "\r\n")

	for _, rowVal := range lines {

		lineMatrix = append(lineMatrix, strings.Split(rowVal, ""))
	}

	return lineMatrix
}

func isInBound(slice [][]string, cord []int) bool {

	if cord[0] < len(slice) && cord[0] >= 0 && cord[1] < len(slice[0]) && cord[1] >= 0 {

		return true
	}

	return false
}

func appendToList(cord []int, locs *[][]int, slice [][]string, table map[string]bool) {
	key := fmt.Sprintf("%d:%d", cord[0], cord[1])
	if isInBound(slice, cord) && !table[key] {
		*locs = append(*locs, cord)
		table[key] = true
	}
}

func getGoodLocsCount(strSlice [][]string, char string, table map[string]bool) [][]int {
	charactors := [][]int{}
	goodLocs := [][]int{}
	for lineIndex, line := range strSlice {
		for el := range line {

			if strSlice[lineIndex][el] == char {
				charactors = append(charactors, []int{lineIndex, el})
				appendToList([]int{lineIndex, el}, &goodLocs, strSlice, table)
			}
		}
	}
	for i := 0; i < len(charactors)-1; i++ {

		for j := i + 1; j < len(charactors); j++ {
			generatePaths(charactors[i], charactors[j], &goodLocs, strSlice, table)
			generatePaths(charactors[j], charactors[i], &goodLocs, strSlice, table)
		}

	}

	return goodLocs
}

func generatePaths(start, end []int, goodLocs *[][]int, slice [][]string, table map[string]bool) {
	// ABS MEANS math.abs()
	rowABS := start[0] - end[0]
	colABS := start[1] - end[1]

	for k := 0; k < UNREACHABLE_NUMBER; k++ {
		// adding the needed amount of differences
		cord := []int{(rowABS + start[0]) + (rowABS * k), (colABS + start[1]) + (colABS * k)}

		appendToList(cord, goodLocs, slice, table)

		if !isInBound(slice, cord) {
			return
		}

	}

}

func getAllChars(slice [][]string) []string {
	charSet := make(map[string]bool)
	chars := []string{}

	for row := range slice {
		for col := range slice[row] {
			if slice[row][col] == "." {
				continue
			}
			if !charSet[slice[row][col]] {
				chars = append(chars, slice[row][col])
			}
			charSet[slice[row][col]] = true
		}
	}

	return chars
}
func answer(path string) int {
	table := make(map[string]bool)
	strSlice := parseStr(path)
	var goodLocs [][]int

	for _, char := range getAllChars(strSlice) {
		goodLocs = append(goodLocs, getGoodLocsCount(strSlice, char, table)...)

	}

	return len(goodLocs)

}
func main() {

	path := "sample"
	if os.Args[1:] != nil && os.Args[1:][0] != "" {
		path = os.Args[1:][0]
	}
	fmt.Println(answer(path))
}
