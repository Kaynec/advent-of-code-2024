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

// func isInBound(matrix [][]string, row, col int) bool {
// 	return row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0])
// }

func moveAndMoveXd(matrix *[][]string, row, col *int, yStep, xStep int) {
	newMatrix := *matrix

	if newMatrix[*row+yStep][*col+xStep] == "#" {
		return
	}

	if newMatrix[*row+yStep][*col+xStep] == "O" {
		curr := newMatrix[*row+yStep][*col+xStep]
		res := [][]int{}
		for i := 1; curr != "#" && curr != "."; i++ {
			res = append(res, []int{*row + yStep*i, *col + xStep*i})
			curr = newMatrix[*row+yStep*i][*col+xStep*i]
		}
		if curr == "#" {
			return
		}
		for i := range res {
			newMatrix[res[i][0]][res[i][1]] = "."
			res[i][0] = *row + (yStep * (i + 1))
			res[i][1] = *col + (xStep * (i + 1))
			newMatrix[res[i][0]][res[i][1]] = "O"
		}
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

		if move == "<" {
			moveAndMoveXd(&matrix, &robotY, &robotX, 0, -1)
		}
		if move == "v" {
			moveAndMoveXd(&matrix, &robotY, &robotX, +1, 0)
		}
		if move == ">" {
			moveAndMoveXd(&matrix, &robotY, &robotX, 0, +1)
		}
		if move == "^" {
			moveAndMoveXd(&matrix, &robotY, &robotX, -1, 0)
		}
	}
	for row := range matrix {
		fmt.Println(matrix[row])
	}
	fmt.Println("\r\n\r\n")
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] != "O" {
				continue
			}
			fmt.Println()
			total += 100*(row) + (col)
		}
	}
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
