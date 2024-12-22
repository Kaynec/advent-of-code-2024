package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func insertSorted(pq []Node, node Node) []Node {
	// Create a new slice to store the result, with enough capacity for the new node
	result := make([]Node, 0, len(pq)+1)

	// Copy elements from pq to result until we find the correct position
	inserted := false
	for _, el := range pq {
		if !inserted && el.cost >= node.cost {
			result = append(result, node)
			inserted = true
		}
		result = append(result, el)
	}

	// If the node has not been inserted, it means it should be appended at the end
	if !inserted {
		result = append(result, node)
	}

	return result
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

type Node struct {
	row, col, rdir, cdir, cost int
	parent                     *Node
}

func getPointFromCord(matrix [][]string, row, col int) int {
	t := map[string]bool{}
	goodPath := map[string]bool{}
	response := math.MaxInt
	pq := []Node{}
	pq = append(pq, Node{row: row, col: col, rdir: 0, cdir: 1, cost: 0})
	for i := 0; len(pq) > 0; i++ {
		popped := pq[0]
		pq = pq[1:]

		t[fmt.Sprintf("%d,%d,%d,%d", popped.row, popped.col, popped.rdir, popped.cdir)] = true

		if matrix[popped.row][popped.col] == "E" {
			curr := popped

			if popped.cost > response {
				break
			}
			for curr.parent != nil {
				goodPath[fmt.Sprintf("%d,%d", curr.row, curr.col)] = true
				curr = *curr.parent
			}
			response = popped.cost
		}

		dirs := []Node{
			{row: popped.row, col: popped.col, rdir: popped.rdir, cdir: popped.cdir, cost: popped.cost + 1},
			{row: popped.row, col: popped.col, rdir: popped.cdir, cdir: popped.rdir, cost: popped.cost + 1001},
			{row: popped.row, col: popped.col, rdir: -popped.cdir, cdir: -popped.rdir, cost: popped.cost + 1001},
		}
		for _, dir := range dirs {
			dir.row, dir.col = dir.row+dir.rdir, dir.col+dir.cdir
			if matrix[dir.row][dir.col] == "#" {
				continue
			}

			key := fmt.Sprintf("%d,%d,%d,%d", dir.row, dir.col, dir.rdir, dir.cdir)
			if t[key] {

			}
			if !t[key] {
				pq = insertSorted(pq, Node{row: dir.row, col: dir.col, rdir: dir.rdir, cdir: dir.cdir, cost: dir.cost, parent: &popped})
			}

		}

	}
	fmt.Println(len(goodPath) + 1)

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
