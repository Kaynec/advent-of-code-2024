package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

)

func print3d[T any](matrix [][]T) {
	for row := range matrix {
		fmt.Println(matrix[row])
	}
}
func stringToInt(char string) int {
	val, _ := strconv.Atoi(char)
	return val
}
func parseStr(path string) [][]any {
	data, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := strings.TrimSpace(string(data))
	matrix := [][]any{}
	for index, line := range strings.Split(str, "\r\n") {
		matrix = append(matrix, []any{})
		for _, col := range strings.Split(line, "") {
			if col == "." {
				matrix[index] = append(matrix[index], 0)
			} else {
				matrix[index] = append(matrix[index], col)
			}
		}
	}
	return matrix
}
func isInBound(matrix [][]any, row, col int) bool {
	return row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0])
}
func paintPoints(matrix [][]any, row, col, curr int) {
	if !isInBound(matrix, row, col) {
		return
	}
	if matrix[row][col] == "#" {
		return
	}

	if matrix[row][col] == 0 {
		matrix[row][col] = curr
	}
}

func pointMatrix(matrix [][]any, row, col, curr int, table map[string]matrixElement) {

	for matrix[row][col] != "S" {
		row = row - 1
		paintPoints(matrix, row-1, col, curr)
		row = row + 1
		paintPoints(matrix, row+1, col, curr)
		col = col + 1
		paintPoints(matrix, row, col+1, curr)
		col = col - 1
		paintPoints(matrix, row, col-1, curr)
		curr++
	}
	fmt.Println("XD XD")
}

//	func isInBound(matrix [][]string, row, col int) bool {
//		return row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0])
//	}
// func getPointFromCord(matrix [][]string, lastVal string, row, col int, count *int, allCount *int, table map[string]bool) {
// 	key := fmt.Sprintf("%d,%d", row, col)

// 	if !isInBound(matrix, row, col) {
// 		*allCount += 1

// 		return
// 	}
// 	if matrix[row][col] != lastVal {
// 		*allCount += 1

// 		return
// 	}
// 	if table[key] {
// 		return
// 	}

// 	*count += 1
// 	tmp := matrix[row][col]
// 	table[key] = true

// 	getPointFromCord(matrix, tmp, row+1, col, count, allCount, table)
// 	getPointFromCord(matrix, tmp, row-1, col, count, allCount, table)
// 	getPointFromCord(matrix, tmp, row, col+1, count, allCount, table)
// 	getPointFromCord(matrix, tmp, row, col-1, count, allCount, table)

// }

type matrixElement struct {
	Row    int
	Col    int
	Number int
}

func getAnswer(path string, times int) int {
	matrix := parseStr(path)

	total := 0
	// 1,13
	// print3d(matrix)
	table := map[string]matrixElement{}
	pointMatrix(matrix, 1, 13, 0, table)
	print3d(matrix)
	// for _, val := range table {

	// 	fmt.Println(val.Row, val.Col, val.Number)
	// }
	// table := make(map[string]bool)
	// for row := range matrix {
	// 	for col := range matrix[row] {
	// 		count := 0
	// 		allCount := 0
	// 		if table[fmt.Sprintf("%d,%d", row, col)] {
	// 			continue
	// 		}
	// 		getPointFromCord(matrix, matrix[row][col], row, col, &count, &allCount, table)
	// 		total += count * allCount
	// 	}
	// }
	// fmt.Println(total)

	return total
}
func getStringFromNum(char string) string {
	num, _ := strconv.Atoi(char)
	return fmt.Sprintf("%d", num)
}
func main() {
	path := "sample"
	times := 25

	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	if len(os.Args) > 2 {
		times = stringToInt(os.Args[2])
	}

	result := getAnswer(path, times)

	fmt.Print(result)
}
