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
	return row >= len(slice) || row < 0 || col >= len(slice[0]) || col < 0
}

func walk(slice [][]string, row, col int) bool {

	dirs := [][]int{{-1, 0}, {0, +1}, {+1, 0}, {0, -1}}
	rowDir, colDir, index := -1, 0, 0

	i := 0
	for {
		i++
		if i > len(slice)*len(slice[0]) {
			return false
		}
		if isEnd(slice, row+rowDir, col+colDir) {
			return true
		}
		if slice[row+rowDir][col+colDir] == "#" {
			if index == len(dirs)-1 {
				index = 0
			} else {
				index += 1
			}

			rowDir, colDir = dirs[index][0], dirs[index][1]

		} else {
			row += rowDir
			col += colDir
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

func partOneAnswer(path string) (total int) {
	sliced := parseStr(path)
	row, col := findCord(sliced)

	for r := range sliced {
		for c := range sliced[r] {
			if sliced[r][c] == "^" {
				continue
			}
			newSliced := make([][]string, len(sliced))
			for i := range sliced {
				newSliced[i] = make([]string, len(sliced[i]))
				copy(newSliced[i], sliced[i]) // Copy each row
			}
			newSliced[r][c] = "#"
			res := walk(newSliced, row, col)
			if !res {
				total++
			}
		}

	}
	return total

}

func main() {
	total := partOneAnswer("input")
	fmt.Println(total)
}
