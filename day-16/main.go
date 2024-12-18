package main

import (
	"fmt"
	"os"
	"sort"
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
func parseStr(path string) [][]string {
	data, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := strings.TrimSpace(string(data))
	matrix := [][]string{}
	for _, line := range strings.Split(str, "\r\n") {
		matrix = append(matrix, strings.Split(line, ""))
	}
	return matrix
}

// func isInBound(matrix [][]any, row, col int) bool {
// 	return row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0])
// }

// func checkAdjacent(matrix [][]any, queue *[][]int, row, col int) {
// 	paint(matrix, queue, row+1, col, matrix[row][col])
// 	paint(matrix, queue, row-1, col, matrix[row][col])
// 	paint(matrix, queue, row, col+1, matrix[row][col])
// 	paint(matrix, queue, row, col-1, matrix[row][col])
// }
// func paint(matrix [][]any, queue *[][]int, row, col int, oldValue any) {
// 	if matrix[row][col] == "E" || matrix[row][col] == "S" || matrix[row][col] == "#" {
// 		return
// 	}
// 	if matrix[row][col] == 0 {
// 		val, ok := oldValue.(int)
// 		if !ok {
// 			val = 0
// 		}
// 		matrix[row][col] = val + 1
// 		*queue = append(*queue, []int{row, col})
// 	}
// }
// func pointMatrix(matrix [][]any, row, col int) {

// 	queue := [][]int{{row, col}}

// 	for len(queue) > 0 {
// 		shift := queue[0]
// 		newRow, newCol := shift[0], shift[1]
// 		queue = queue[1:]
// 		checkAdjacent(matrix, &queue, newRow, newCol)
// 	}
// }

func isInBound(matrix [][]string, row, col int) bool {
	return row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0])
}
func getPointFromCord(matrix [][]string, row, col int, table map[string]bool, res [][]int, count, allCount *int, dir string) {

	key := fmt.Sprintf("%d,%d", row, col)
	if !isInBound(matrix, row, col) || matrix[row][col] == "#" {
		return
	}
	if matrix[row][col] == "E" {
		sort.Slice(res, func(i, j int) bool {
			return res[i][0] < res[j][0]
		})
		fmt.Println(*allCount, len(res))
		return
	}

	if table[key] {
		return
	}

	// if changeDirection {
	// 	*count += 1
	// } else {
	// }
	*allCount += 1

	table[key] = true
	res = append(res, []int{row, col})

	// oldDir := dir

	dirs := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	for _, direction := range dirs {
		dirRow, dirCol := direction[0], direction[1]
		getPointFromCord(matrix, row+dirRow, col+dirCol, table, res, count, allCount, dir)

		// 	rowDiff, colDiff := math.Abs(float64(row+dirRow)), math.Abs(float64(col+dirCol))
		// 	if rowDiff != 0 {
		// 		dir = "Y"
		// 	} else if colDiff != 0 {
		// 		dir = "X"
		// 	}
		// 	if oldDir != dir {
		// 		getPointFromCord(matrix, row+dirRow, col+dirCol, table, res, count, allCount, dir)
		// 	} else {
		// 		getPointFromCord(matrix, row+dirRow, col+dirCol, table, res, count, allCount, dir)
		// 	}

	}

}

func getAnswer(path string) int {
	matrix := parseStr(path)

	total := 0

	print3d(matrix)

	dirs := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	row, col := 12, 1
	for range dirs {
		res := [][]int{}
		dir := ""
		table := map[string]bool{}
		count, allCount := 0, 0
		getPointFromCord(matrix, row, col, table, res, &count, &allCount, dir)
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
