package main

import (
	"fmt"
	"os"
	"strings"
)

func parseStr(path string) [][]string {
	data, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := strings.TrimSpace(string(data))
	matrix := [][]string{}
	for _, line := range strings.Split(str, "\r\n") {
		matrix = append(matrix, strings.Split(line, ""))
	}
	return matrix
}

type Node struct {
	row, col, rdir, cdir, chances int
	parent                        *Node
}

func findPath(matrix [][]string, row, col int, isFirstTime bool) (int, [][]int) {
	t := map[string]bool{}
	pq := []Node{}
	pq = append(pq, Node{row: row, col: col, rdir: 0, cdir: 1, chances: 2})
	goodPath := map[string]bool{}

	defaultPoints := [][]int{}

	for i := 0; len(pq) > 0; i++ {
		popped := pq[0]
		pq = pq[1:]

		t[fmt.Sprintf("%d,%d,%d,%d", popped.row, popped.col, popped.rdir, popped.cdir)] = true

		if matrix[popped.row][popped.col] == "E" {

			curr := popped

			for curr.parent != nil {
				goodPath[fmt.Sprintf("%d,%d", curr.row, curr.col)] = true
				curr = *curr.parent
			}
			return len(goodPath), defaultPoints
		}

		dirs := []Node{
			{row: popped.row, col: popped.col, rdir: 1, cdir: 0},
			{row: popped.row, col: popped.col, rdir: -1, cdir: 0},
			{row: popped.row, col: popped.col, rdir: 0, cdir: 1},
			{row: popped.row, col: popped.col, rdir: 0, cdir: -1},
		}
		for _, dir := range dirs {

			if matrix[dir.row+dir.rdir][dir.col+dir.cdir] == "#" {
				currRow, currCol := dir.row+dir.rdir, dir.col+dir.cdir
				if isFirstTime {
					if currRow == 0 || currRow >= len(matrix)-1 || currCol == 0 || currCol >= len(matrix[0])-1 {
						continue
					}
					if !containsIntSlice(defaultPoints, []int{currRow, currCol, dir.row + dir.rdir*1, dir.col + (dir.cdir * 1)}) {
						defaultPoints = append(defaultPoints, []int{currRow, currCol, dir.row + (dir.rdir * 1), dir.col + (dir.cdir * 1)})
					}

				}
				continue
			}
			dir.row = dir.row + dir.rdir
			dir.col = dir.col + dir.cdir

			key := fmt.Sprintf("%d,%d,%d,%d", dir.row, dir.col, dir.rdir, dir.cdir)
			if t[key] {

			}
			if !t[key] {
				pq = append(pq, Node{row: dir.row, col: dir.col, rdir: dir.rdir, cdir: dir.cdir, parent: &popped})
			}

		}

	}
	return len(goodPath), defaultPoints
}

func getPointFromCord(matrix [][]string, row, col int) int {

	initialCount, defaultPoints := findPath(matrix, row, col, true)

	mapping := map[int]int{}
	for i, point := range defaultPoints {
		currPoint, nextPoint := []int{point[0], point[1]}, []int{point[2], point[3]}

		currTmp, nextTmp := matrix[currPoint[0]][currPoint[1]], matrix[nextPoint[0]][nextPoint[1]]

		if currTmp == "." || nextTmp == "." {
			continue
		}

		matrix[currPoint[0]][currPoint[1]] = "."
		matrix[nextPoint[0]][nextPoint[1]] = "."

		count, _ := findPath(matrix, row, col, false)

		matrix[currPoint[0]][currPoint[1]] = currTmp
		matrix[nextPoint[0]][nextPoint[1]] = nextTmp
		if count == 0 || count == initialCount {
			continue
		}

		fmt.Println(i)

		mapping[initialCount-count] += 1
	}

	total := 0
	for key, value := range mapping {
		if key >= 100 {
			total += value
		}
	}

	fmt.Println(mapping, total)

	return 0
}

func getAnswer(path string) int {
	matrix := parseStr(path)
	row, col := 0, 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == "S" {
				row, col = i, j
			}
		}
	}
	getPointFromCord(matrix, row, col)
	return 0
}

func main() {
	path := "sample"

	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	result := getAnswer(path)

	fmt.Println(result, "result")
}

func containsIntSlice(source [][]int, element []int) bool {
	for _, value := range source {
		if value[0] == element[0] && value[1] == element[1] && value[2] == element[2] && value[3] == element[3] {
			return true
		}
	}
	return false
}
