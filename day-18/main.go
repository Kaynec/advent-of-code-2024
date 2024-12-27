package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ROW         = 70
	COL         = 70
	CURRPTBYTES = 1024
)

func stringToInt(char string) int {
	val, _ := strconv.Atoi(char)
	return val
}

func makeMatrix() [][]string {

	matrix := [][]string{}
	for row := range ROW + 1 {
		matrix = append(matrix, []string{})
		for range COL + 1 {
			matrix[row] = append(matrix[row], ".")
		}
	}
	return matrix
}

func parseStr(path string) [][]int {
	data, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := strings.TrimSpace(string(data))
	vertices := [][]int{}
	for _, line := range strings.Split(str, "\r\n") {
		fields := strings.Split(line, ",")
		x, y := stringToInt(fields[0]), stringToInt(fields[1])
		vertices = append(vertices, []int{y, x})
	}
	return vertices
}

func isInBound(matrix [][]string, row, col int) bool {
	return row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0])
}

func getPointFromCord(matrix [][]string, row, col int) (int, bool) {
	t := map[string]bool{}

	queue := [][]int{{row, col, 1}}

	dirs := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	for i := 0; len(queue) > 0; i++ {
		popped := queue[0]
		queue = queue[1:]

		row, col, step := popped[0], popped[1], popped[2]
		for _, dir := range dirs {
			row, col := row+dir[0], col+dir[1]
			if !isInBound(matrix, row, col) {
				continue
			}
			if matrix[row][col] == "#" {
				continue
			}
			if row == ROW && col == COL {
				return step, true
			}
			key := fmt.Sprintf("%d,%d", row, col)
			if !t[key] {
				queue = append(queue, []int{row, col, step + 1})
			}

			t[key] = true
		}

	}

	return 0, false
}

func getAnswer(path string) (int, int) {
	matrix := makeMatrix()
	vertices := parseStr(path)
	for _, vertex := range vertices[CURRPTBYTES:] {
		row, col := vertex[0], vertex[1]
		matrix[row][col] = "#"
		_, ok := getPointFromCord(matrix, 0, 0)
		if !ok {
			return vertex[1], vertex[0]
		}
	}
	return 0, 0
}

func main() {
	path := "sample"

	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	x, y := getAnswer(path)

	fmt.Print(x, y)
}
