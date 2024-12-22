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
	for _, line := range strings.Split(str, "\r\rdir") {
		matrix = append(matrix, strings.Split(line, ""))
	}
	return matrix
}

type Node struct {
	row, col, rdir, cdir, cost int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() *Node {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

func getPointFromCord(matrix [][]string, row, col int) (int, bool) {
	t := map[string]bool{}

	pq := &PriorityQueue{}
	pq.Push(&Node{
		row:  row,
		col:  col,
		rdir: 0,
		cdir: 1,
		cost: 0,
	})
	fmt.Println(row, col)
	for i := 0; pq.Len() > 0; i++ {
		popped := pq.Pop()
		row, col, rDir, cDir, cost := popped.row, popped.col, popped.rdir, popped.cdir, popped.cost

		if matrix[row][col] == "E" {
			fmt.Println("Xo Xo", row, col, rDir, cDir, cost)
		}
		fmt.Println(row, col)

		dirs := [][]int{
			{row, col, rDir, cDir, cost + 1},
			{row, col, cDir, rDir, cost + 1},
			{row, col, -cDir, -rDir, cost + 1},
		}
		for _, dir := range dirs {
			row, col := row+dir[2], col+dir[3]
			if matrix[row][col] == "#" || matrix[row][col] == "S" {
				continue
			}

			if key := fmt.Sprintf("%d,%d", row, col); !t[key] {
				pq.Push(&Node{
					row:  dir[0],
					col:  dir[1],
					rdir: dir[2],
					cdir: dir[3],
					cost: dir[4],
				})
			}

		}

		t[fmt.Sprintf("%d,%d", row, col)] = true

	}

	return 0, true
}

func getAnswer(path string) int {
	matrix := parseStr(path)
	// row, col := 0, 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == "S" {
				// row, col = i, j
				fmt.Println(matrix[i][j])
				fmt.Println(i)
			}
		}
	}

	// fmt.Println(row)
	// getPointFromCord(matrix, row, col)
	return 0
}

func main() {
	path := "sample"

	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	result := getAnswer(path)

	fmt.Print(result)
}
