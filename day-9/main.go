package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func toNum(char string) int {
	val, _ := strconv.Atoi(char)
	return val
}

func findAllSubstring(str []any, char any) [][]int {

	res := [][]int{}

	for i := 0; i < len(str)-1; i++ {
		if str[i] != char {
			continue
		}

		for j := i; j < len(str); j++ {
			if str[j] == char && j != len(str)-1 {
				continue
			}
			if j == len(str)-1 {
				res = append(res, []int{i, j + 1})
				i = j - 1
			} else {
				res = append(res, []int{i, j})
				i = j
			}
			break
		}

	}

	return res

}

func parseStr(path string) [][]int {
	val, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := strings.Split(string(val), "")

	var lineMatrix [][]int

	for i := 0; i < len(str); i += 2 {
		if i+1 == len(str) {
			lineMatrix = append(lineMatrix, []int{toNum(str[i]), 0})
			break
		}

		lineMatrix = append(lineMatrix, []int{toNum(str[i]), toNum(str[i+1])})

	}

	return lineMatrix
}

func getAnswer(path string) int {
	matrix := parseStr(path)

	str := extrctSliceFromTupleInt(matrix)
	eCord := findAllSubstring(str, ".")

	normalCords := getNormalCards(str)

	formFinalResult(normalCords, str, eCord)

	total := 0

	for multiplier, value := range str {
		if value == "." {
			continue
		}
		if s, ok := value.(int); ok {
			total += multiplier * s
		}
	}

	return total

}

func formFinalResult(normalCords [][]int, str []any, eCord [][]int) {
	for _, cord := range normalCords {

		cordItem := str[cord[0]]

		for index, e := range eCord {
			len := e[1] - e[0]
			diff := cord[1] - cord[0]
			if diff > len || cord[0] < e[0] {
				continue
			}
			for i := 0; i < diff; i++ {
				str[i+e[0]] = cordItem
				str[i+cord[0]] = "."
			}

			if len == diff {
				// remove the filled dots slice
				eCord = append(eCord[0:index], eCord[index+1:]...)
			} else {
				// or modify it if doesn't fit
				e[0] = e[0] + cord[1] - cord[0]
			}
			break

		}
	}
}

func extrctSliceFromTupleInt(matrix [][]int) []any {
	str := []any{}

	for id, slice := range matrix {

		for i := 0; i < slice[0]; i++ {

			str = append(str, id)
		}
		for i := 0; i < slice[1]; i++ {

			str = append(str, ".")
		}
	}
	return str
}

func getNormalCards(str []any) [][]int {
	table := make(map[int]bool)

	var normalCords [][]int
	for i := range str {
		if num, ok := str[i].(int); ok {
			if table[num] {
				continue
			}
			normalCords = append(normalCords, findAllSubstring(str, num)...)

			table[num] = true
		}
	}
	slices.Reverse(normalCords)
	return normalCords
}
func main() {
	path := "sample"
	if os.Args[1:] != nil && os.Args[1:][0] != "" {
		path = os.Args[1:][0]
	}
	result := getAnswer(path)
	fmt.Println(result)
}
