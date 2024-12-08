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

func isEnd(slice [][]string, row, col int) bool {
	fmt.Println(row, col)
	if slice[row][col] == "#" || row >= len(slice) || row < 0 || col >= len(slice[0]) || col < 0 {
		fmt.Println("REACHED END", row, col)
		return true
	}
	return false
}

func filter(res [][]int) [][]int {
	newSlice := [][]int{}
	table := make(map[string]bool)

	for _, slice := range res {
		row, col := slice[0], slice[1]

		if table[fmt.Sprintf("%d,%d", row, col)] {
			continue
		}
		newSlice = append(newSlice, []int{row, col})
		table[fmt.Sprintf("%d,%d", row, col)] = true
	}
	return newSlice
}

func walk(slice [][]string, row, col int) [][]int {

	dir := "up"
	res := [][]int{}

	i := 0
	for {
		i++
		res = append(res, []int{row, col})

		if dir == "up" {
			if row-1 < 0 || slice[row-1][col] == "#" {
				if isEnd(slice, row, col+1) {
					return res
				}
				dir = "right"
			} else {
				row = row - 1
			}
		}
		if dir == "bottom" {
			if row+1 >= len(slice) || slice[row+1][col] == "#" {
				if isEnd(slice, row, col-1) {
					return res
				}
				dir = "left"
			} else {
				row = row + 1
			}
		}

		if dir == "right" {
			if col+1 >= len(slice[0]) || slice[row][col+1] == "#" {
				if isEnd(slice, row+1, col) {
					return res
				}
				dir = "bottom"
			} else {
				col = col + 1
			}
		}
		if dir == "left" {

			if col-1 < 0 || slice[row][col-1] == "#" {
				if isEnd(slice, row-1, col) {
					return res
				}
				dir = "up"
			} else {
				col = col - 1
			}
		}
	}
}

func findCord(sliced [][]string) (int, int) {

	for rowIdx, rowValue := range sliced {
		for colValue := range rowValue {
			if sliced[rowIdx][colValue] == "^" {

				return rowIdx, colValue
			}
		}
	}
	return 0, 0
}

func partOneAnswer(path string) {
	sliced := parseStr(path)
	row, col := findCord(sliced)

	res := walk(sliced, row, col)

	fmt.Println(len(filter(res)))
}

func main() {
	partOneAnswer("input")
}
