package main

import (
	"fmt"
	"os"
	"strings"
)

func parseStr(path string) ([][]string, string) {
	data, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := strings.TrimSpace(string(data))

	matrix := [][]string{}

	matrixStr := strings.Split(str, "\r\n\r\n")[0]
	moves := strings.Split(str, "\r\n\r\n")[1]
	for _, str := range strings.Split(matrixStr, "\r\n") {
		matrix = append(matrix, strings.Split(str, ""))
	}

	return matrix, moves
}

func scaledWareHouse(matrix [][]string) [][]string {

	mapping := map[string][]string{
		"#": {"#", "#"},
		"O": {"[", "]"},
		".": {".", "."},
		"@": {"@", "."},
	}

	newMatrix := [][]string{}
	for row := range matrix {
		newMatrix = append(newMatrix, []string{})
		for col := range matrix[row] {
			val := mapping[matrix[row][col]]
			newMatrix[row] = append(newMatrix[row], val...)
		}
	}
	return newMatrix
}

func getConnectedOnYAxis(matrix [][]string, row, col int, yStep int, res *[][]int, shouldMove *bool) {
	if matrix[row][col] == "#" {
		*shouldMove = false
		return
	}

	if matrix[row][col] == "." {
		return
	}
	charMap := map[string]int{
		"[": +1,
		"]": -1,
	}
	val := charMap[matrix[row][col]]

	*res = append(*res, []int{row, col})
	*res = append(*res, []int{row, col + val})
	getConnectedOnYAxis(matrix, row+yStep, col, yStep, res, shouldMove)
	getConnectedOnYAxis(matrix, row+yStep, col+val, yStep, res, shouldMove)
}

func moveAndMoveXd(matrix *[][]string, row, col *int, yStep, xStep int) {
	newMatrix := *matrix

	res := [][]int{}

	if newMatrix[*row+yStep][*col+xStep] == "#" {
		return
	}

	// handle y movement
	if yStep != 0 {
		shouldMove := true
		getConnectedOnYAxis(*matrix, *row+yStep, *col+xStep, yStep, &res, &shouldMove)
		if !shouldMove {
			return
		}

	}

	// handle x movement
	if newMatrix[*row+yStep][*col+xStep] == "[" || newMatrix[*row+yStep][*col+xStep] == "]" {
		curr := newMatrix[*row+yStep][*col+xStep]
		for i := 1; curr != "#" && curr != "."; i++ {
			curr = newMatrix[*row+yStep*i][*col+xStep*(i)]
			if curr != "." && curr != "#" {
				res = append(res, []int{*row + yStep, *col + xStep*i})
			}
		}
		if curr == "#" {
			return
		}
	}
	shapes := []string{}
	for _, el := range res {
		shapes = append(shapes, newMatrix[el[0]][el[1]])
	}
	for _, el := range res {
		newMatrix[el[0]][el[1]] = "."
	}
	for i, el := range res {
		el[0] = el[0] + yStep
		el[1] = el[1] + xStep
		newMatrix[el[0]][el[1]] = shapes[i]
	}

	newMatrix[*row][*col] = "."
	*row = *row + yStep
	*col = *col + xStep
	newMatrix[*row][*col] = "@"

	*matrix = newMatrix
}

func getAnswer(path string) int {
	total := 0
	robotX := 0
	robotY := 0
	matrix, moves := parseStr(path)
	matrix = scaledWareHouse(matrix)

	directions := map[string][]int{
		"<": {0, -1},
		">": {0, +1},
		"^": {-1, 0},
		"v": {+1, 0},
	}

	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == "@" {
				robotX = col
				robotY = row
				break

			}
		}
	}

	matrix[robotY][robotX] = "."
	for _, move := range strings.Split(moves, "") {

		dir, ok := directions[move]
		if !ok {
			continue
		}
		moveAndMoveXd(&matrix, &robotY, &robotX, dir[0], dir[1])
	}
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] != "[" {
				continue
			}
			total += 100*(row) + (col)
		}
	}
	return total
}

func main() {
	path := "sample"

	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	result := getAnswer(path)

	fmt.Print(result)
}
