package main

import (
	"fmt"
	"os"
	"sort"
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
	row, col, rdir, cdir int
	parent               *Node
}

func getPointFromCord(matrix [][]string, row, col int) int {
	t := map[string]bool{}
	goodPath := map[string]bool{}
	defaultPoints := [][]int{}
	pq := []Node{}

	pq = append(pq, Node{row: row, col: col, rdir: 0, cdir: 1})
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
		}

		dirs := []Node{
			{row: popped.row + 1, col: popped.col},
			{row: popped.row - 1, col: popped.col},
			{row: popped.row, col: popped.col + 1},
			{row: popped.row, col: popped.col - 1},
		}
		for _, dir := range dirs {
			if matrix[dir.row][dir.col] == "#" {
				if !containtsIntSlice(defaultPoints, []int{dir.row, dir.col}) {
					defaultPoints = append(defaultPoints, []int{dir.row, dir.col})
				}

				continue
			}

			key := fmt.Sprintf("%d,%d,%d,%d", dir.row, dir.col, dir.rdir, dir.cdir)
			if t[key] {

			}
			if !t[key] {
				pq = append(pq, Node{row: dir.row, col: dir.col, rdir: dir.rdir, cdir: dir.cdir, parent: &popped})
			}

		}

	}
	sort.Slice(defaultPoints, func(i, j int) bool {
		if defaultPoints[i][0] == defaultPoints[j][0] {
			return defaultPoints[i][1] < defaultPoints[j][1]

		}
		return defaultPoints[i][0] < defaultPoints[j][0]
	})
	fmt.Println(len(goodPath), len(defaultPoints), defaultPoints)
	// fmt.Println(defaultPoints)

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

func containtsIntSlice(source [][]int, element []int) bool {
	for _, value := range source {
		if value[0] == element[0] && value[1] == element[1] {
			return true
		}
	}
	return false
}
