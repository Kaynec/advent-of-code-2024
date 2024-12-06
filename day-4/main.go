package main

import (
	"fmt"
	"os"
	"slices"
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

func checkAnswer(slice [][]string, nums [][]int, answer string, shouldSort bool) bool {

	s := ""
	for _, num := range nums {
		row, col := num[0], num[1]

		if row >= len(slice) || row < 0 || col >= len(slice[0]) || col < 0 {
			return false
		}
		s += slice[row][col]
	}

	if !shouldSort {
		return s == answer
	}

	newS := strings.Split(s, "")

	slices.Sort(newS)

	newAnswer := strings.Split(answer, "")

	slices.Sort(newAnswer)

	return strings.Join(newS, "") == strings.Join(newAnswer, "")
}

func checkXmas(slice [][]string, row int, col int, numSlice [][][]int, answer string, shouldSort bool) int {

	sum := 0

	for _, matrix := range numSlice {

		if checkAnswer(slice, matrix, answer, shouldSort) {
			sum++
		}
	}

	return sum
}

func partOneAnswer(path string) int {
	mulSlice := parseStr(path)

	sum := 0

	for row := range mulSlice {

		for col := range mulSlice[row] {

			templateSlice := [][][]int{
				[][]int{[]int{row, col}, []int{row, col + 1}, []int{row, col + 2}, []int{row, col + 3}},
				[][]int{[]int{row, col}, []int{row, col - 1}, []int{row, col - 2}, []int{row, col - 3}},
				[][]int{[]int{row, col}, []int{row - 1, col}, []int{row - 2, col}, []int{row - 3, col}},
				[][]int{[]int{row, col}, []int{row + 1, col}, []int{row + 2, col}, []int{row + 3, col}},
				[][]int{[]int{row, col}, []int{row - 1, col - 1}, []int{row - 2, col - 2}, []int{row - 3, col - 3}},
				[][]int{[]int{row, col}, []int{row - 1, col + 1}, []int{row - 2, col + 2}, []int{row - 3, col + 3}},
				[][]int{[]int{row, col}, []int{row + 1, col - 1}, []int{row + 2, col - 2}, []int{row + 3, col - 3}},
				[][]int{[]int{row, col}, []int{row + 1, col + 1}, []int{row + 2, col + 2}, []int{row + 3, col + 3}},
			}

			if mulSlice[row][col] == "X" {
				sum += checkXmas(mulSlice, row, col, templateSlice, "XMAS", false)
			}

		}
	}

	fmt.Println("SUM", sum)
	return sum

}

func partTwoAnswer(path string) {
	mulSlice := parseStr(path)

	sum, masTreshhold := 0, 2

	for row := range mulSlice {

		for col := range mulSlice[row] {
			if mulSlice[row][col] != "A" {
				continue
			}

			templateSlice := [][][]int{
				[][]int{[]int{row, col}, []int{row - 1, col - 1}, []int{row + 1, col + 1}},
				[][]int{[]int{row, col}, []int{row - 1, col + 1}, []int{row + 1, col - 1}},
			}

			if checkXmas(mulSlice, row, col, templateSlice, "MAS", true) == masTreshhold {

				sum += 1
			}

		}
	}

	fmt.Println("SUM", sum)
}

func main() {
	partOneAnswer("input")
	partTwoAnswer("input")
}
