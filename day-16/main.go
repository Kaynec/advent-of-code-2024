package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

)

func print3d[T any](matrix [][]T) {
	newMatrix := [][]string{}
	for row := range matrix {
		newMatrix = append(newMatrix, []string{})
		for _, col := range matrix[row] {
			val, ok := any(col).(string)
			if ok {
				newMatrix[row] = append(newMatrix[row], fmt.Sprintf("%3s", val))
			}

			num, isOk := any(col).(int)
			if isOk {

				newMatrix[row] = append(newMatrix[row], fmt.Sprintf("%3d", num))
			}
		}
		fmt.Println(newMatrix[row])
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

// func isInBound(matrix [][]any, row, col int) bool {
// 	return row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0])
// }

func checkAdjacent(matrix [][]any, queue *[][]int, row, col int) {
	paint(matrix, queue, row+1, col, matrix[row][col])
	paint(matrix, queue, row-1, col, matrix[row][col])
	paint(matrix, queue, row, col+1, matrix[row][col])
	paint(matrix, queue, row, col-1, matrix[row][col])
}
func paint(matrix [][]any, queue *[][]int, row, col int, oldValue any) {
	if matrix[row][col] == "E" || matrix[row][col] == "S" || matrix[row][col] == "#" {
		return
	}
	if matrix[row][col] == 0 {
		val, ok := oldValue.(int)
		if !ok {
			val = 0
		}
		matrix[row][col] = val + 1
		*queue = append(*queue, []int{row, col})
	}
}
func pointMatrix(matrix [][]any, row, col int) {

	queue := [][]int{{row, col}}

	for len(queue) > 0 {
		shift := queue[0]
		newRow, newCol := shift[0], shift[1]
		queue = queue[1:]
		checkAdjacent(matrix, &queue, newRow, newCol)
	}
}

func isInBound(matrix [][]any, row, col int) bool {
	return row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0])
}
func getPointFromCord(matrix [][]any, row, col, lastVal int, table map[string]bool, count, allCount *int) {

	key := fmt.Sprintf("%d,%d", row, col)
	if !isInBound(matrix, row, col) || matrix[row][col] == "#" {
		return
	}
	if matrix[row][col] == "E" {
		fmt.Println(*count)
		fmt.Println("AT LAST")
	}
	if matrix[row][col] != lastVal-1 {
		return
	}
	if table[key] {
		return
	}

	*count += 1
	tmp := matrix[row][col].(int)
	table[key] = true

	getPointFromCord(matrix, row+1, col, tmp, table, count, allCount)
	getPointFromCord(matrix, row-1, col, tmp, table, count, allCount)
	getPointFromCord(matrix, row, col+1, tmp, table, count, allCount)
	getPointFromCord(matrix, row, col-1, tmp, table, count, allCount)

}

func getAnswer(path string) int {
	matrix := parseStr(path)

	total := 0
	// 1,13
	// print3d(matrix)
	pointMatrix(matrix, 1, 13)
	print3d(matrix)

	dirs := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	row, col := 12, 1
	for _, dir := range dirs {
		table := map[string]bool{}

		val, ok := matrix[row+dir[0]][col+dir[1]].(int)
		count, allCount := 0, 0
		if ok {
			getPointFromCord(matrix, row+dir[0], col+dir[1], val+1, table, &count, &allCount)
		}
	}

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

	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	result := getAnswer(path)

	fmt.Print(result)
}
