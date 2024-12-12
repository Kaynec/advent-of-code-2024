package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseStr(path string) (matrix [][]int, zeros [][]int) {
	val, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := string(val)

	for row, line := range strings.Split(str, "\r\n") {
		matrix = append(matrix, []int{})
		for col, num := range strings.Split(line, "") {
			num, _ := strconv.Atoi(num)
			matrix[row] = append(matrix[row], num)
			if num == 0 {
				zeros = append(zeros, []int{row, col})
			}
		}
	}

	return matrix, zeros
}

func isInBound(matrix [][]int, row, col int) bool {
	return row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0])
}
func walk(matrix [][]int, row, col int, lastVal int, result *[][]int) {

	if !isInBound(matrix, row, col) || matrix[row][col]-lastVal != 1 {
		return
	}
	if matrix[row][col] == 9 {
		*result = append(*result, []int{row, col})
	}
	walk(matrix, row, col-1, matrix[row][col], result)
	walk(matrix, row, col+1, matrix[row][col], result)
	walk(matrix, row-1, col, matrix[row][col], result)
	walk(matrix, row+1, col, matrix[row][col], result)
}
func partOne(path string) int {
	matrix, zeros := parseStr(path)
	total := 0
	for _, zero := range zeros {
		result := [][]int{}
		walk(matrix, zero[0], zero[1], -1, &result)
		total += len(result)
	}
	return total
}
func main() {
	path := "sample"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	fmt.Println(partOne(path))
}
