package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

)

func stringToInt(char string) int {
	val, _ := strconv.Atoi(char)
	return val
}
func parseStr(path string) [][]string {
	data, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := strings.TrimSpace(string(data))
	matrix := [][]string{}
	for index, line := range strings.Split(str, "\r\n") {
		matrix = append(matrix, []string{})
		for _, el := range strings.Split(line, "") {
			matrix[index] = append(matrix[index], el)
		}
	}
	return matrix
}
func isInBound(matrix [][]string, row, col int) bool {
	return row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0])
}
func getPointFromCord(matrix [][]string, lastVal string, row, col int, count *int, allCount *int, table map[string]bool) {
	key := fmt.Sprintf("%d,%d", row, col)

	if !isInBound(matrix, row, col) {
		*allCount += 1

		return
	}
	if matrix[row][col] != lastVal {
		*allCount += 1

		return
	}
	if table[key] {
		return
	}

	*count += 1
	tmp := matrix[row][col]
	table[key] = true

	getPointFromCord(matrix, tmp, row+1, col, count, allCount, table)
	getPointFromCord(matrix, tmp, row-1, col, count, allCount, table)
	getPointFromCord(matrix, tmp, row, col+1, count, allCount, table)
	getPointFromCord(matrix, tmp, row, col-1, count, allCount, table)

}

func getAnswer(path string, times int) int {
	matrix := parseStr(path)
	fmt.Println(matrix)
	total := 0
	table := make(map[string]bool)
	for row := range matrix {
		for col := range matrix[row] {
			count := 0
			allCount := 0
			if table[fmt.Sprintf("%d,%d", row, col)] {
				continue
			}
			getPointFromCord(matrix, matrix[row][col], row, col, &count, &allCount, table)
			total += count * allCount
		}
	}
	fmt.Println(total)

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
