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

// func getNextMove(slice [][]string, row, col int){

// }

func isEnd(slice [][]string, row, col int) bool {
	if slice[row][col] == "#" {
		fmt.Println("REACHED END", row, col)
		return true
	}
	return false
}

func filter(res [][]int) [][]int {
	newSlice := [][]int{}
	table := make(map[string]bool)

	for _, slice := range res {

		if table[fmt.Sprintf("%d,%d", slice[0], slice[1])] {
			continue
		}
		newSlice = append(newSlice, []int{slice[0], slice[1]})
		table[fmt.Sprintf("%d,%d", slice[0], slice[1])] = true
	}
	return newSlice
}

func walk(slice [][]string, row, col int) [][]int {

	dir := "up"
	res := [][]int{}

	i := 0
	for {
		i++
		fmt.Println(i)
		res = append(res, []int{row, col})

		if dir == "up" {
			if row-1 < 0 || slice[row-1][col] == "#" {
				dir = "right"
				if isEnd(slice, row, col+1) {
					return res
				}
			} else {
				row = row - 1
			}
		}
		if dir == "bottom" {
			if row+1 >= len(slice) || slice[row+1][col] == "#" {
				dir = "left"
				if isEnd(slice, row, col-1) {
					return res
				}
			} else {
				row = row + 1
			}
		}

		if dir == "right" {
			if col+1 >= len(slice[0]) || slice[row][col+1] == "#" {
				dir = "bottom"
				if isEnd(slice, row+1, col-1) {
					return res
				}
			} else {
				col = col + 1
			}
		}
		if dir == "left" {
			if col-1 < 0 || slice[row][col-1] == "#" {
				dir = "up"
				if isEnd(slice, row-1, col-1) {
					return res
				}
			} else {
				col = col - 1
			}
		}
	}
}

func partOneAnswer(path string) {

	sliced := parseStr(path)

	col, row := 6, 4
	res := walk(sliced, col, row)

	fmt.Println(len(filter(res)))
}

func main() {
	partOneAnswer("input")
}
