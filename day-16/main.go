package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	pq "gopkg.in/dnaeon/go-priorityqueue.v1"

)

// func printMatrix(matrix [][]string, res [][]int) {
// 	newMatrix := [][]string{}
// 	fmt.Println("\r\n\r\n")
// 	for row := range matrix {
// 		newMatrix = append(newMatrix, []string{})
// 		for col := range matrix[row] {

// 			ok := false
// 			for _, item := range res {
// 				if item[0] == row && item[1] == col {
// 					// index = item
// 					ok = true
// 				}
// 			}
// 			if !ok {
// 				newMatrix[row] = append(newMatrix[row], matrix[row][col])
// 				continue
// 			} else {
// 				newMatrix[row] = append(newMatrix[row], "?")

// 			}

// 		}
// 		fmt.Println(newMatrix[row])
// 	}
// 	fmt.Println("\r\n\r\n")
// }

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
		// fmt.Println(newMatrix[row])
	}
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

func getResult(allResponse [][][]int) int {
	response := 0
	for result := range allResponse {
		count := 0
		total := 2
		direction := []int{0, 1}

		last := allResponse[result][0]
		curr := allResponse[result][1]

		if last[0]-curr[0] != direction[1] || last[1]-curr[1] != direction[1] {
			count += 1
			direction[0], direction[1] = last[0]-curr[0], last[1]-curr[1]
		}

		for i := 2; i < len(allResponse[result]); i++ {
			last = allResponse[result][i-1]
			curr = allResponse[result][i]

			if last[0]-curr[0] != direction[0] || last[1]-curr[1] != direction[1] {
				count += 1
				direction[0], direction[1] = last[0]-curr[0], last[1]-curr[1]
			}

			total += 1
		}

		response = min(response, total+count*1000)
	}
	return response
}

func makeKey(row, col, rDir, cDir int) string {

	return fmt.Sprintf("%d,%d,%d,%d", row, col, rDir, cDir)
}
func getFromKey(key string) (int, int, int, int) {
	row, _ := strconv.Atoi(strings.Split(key, ",")[0])
	col, _ := strconv.Atoi(strings.Split(key, ",")[1])
	rDir, _ := strconv.Atoi(strings.Split(key, ",")[2])
	cDir, _ := strconv.Atoi(strings.Split(key, ",")[3])
	return row, col, rDir, cDir
}

func handleLoop(row, rDir, col, cDir int, matrix [][]string, queue *pq.PriorityQueue[string, int64], t map[string]bool, count int) {

	row, col = row+rDir, col+cDir
	fmt.Println(row, col, "+++++", rDir, cDir)
	if matrix[row][col] == "#" || matrix[row][col] == "S" {
		return
	}
	if matrix[row][col] == "E" {
		fmt.Println("Reached End Xo Xo", row, col)
		return
	}
	key := fmt.Sprintf("%d,%d", row, col)

	if t[key] {
		return
	}

	cost := count

	queue.Put(makeKey(row, col, rDir, cDir), int64(cost))
	t[key] = true
}

func getPointFromCord(matrix [][]string, row, col int) (int, bool) {
	dir := []int{0, 1}
	t := map[string]bool{}
	queue := pq.New[string, int64](pq.MinHeap)
	queue.Put(makeKey(row, col, dir[0], dir[1]), 1)
	count := 0
	fmt.Println("XO XO")

	for i := 0; queue.Len() > 0; i++ {
		popped := queue.Get()
		row, col, rDir, cDir := getFromKey(popped.Value)

		dir[0], dir[1] = rDir, cDir
		count += 1
		handleLoop(row, dir[0], col, dir[1], matrix, queue, t, 1)
		handleLoop(row, dir[1], col, dir[0], matrix, queue, t, 1)
		handleLoop(row, (-dir[1]), col, (-dir[0]), matrix, queue, t, 1)
	}
	fmt.Println(count)
	// fmt.Println(row+dir[0], col+dir[1])
	// fmt.Println(row+dir[1], col+dir[0])
	// fmt.Println(row+-dir[1], col+-dir[0])
	return 0, true
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
	response := math.MaxInt

	// allResponse := [][][]int{}
	// cache := map[string][][]int{}
	// i := 0
	// for range 4 {
	// 	// res := [][]int{}
	// 	table := map[string]bool{}
	// 	getPointFromCord(matrix, row, col, 0, row, col, "right", table)
	// }
	allResponse := [][][]int{}

	getPointFromCord(matrix, row, col)

	// for range 4 {

	// 	getPointFromCord(matrix, row, col)
	// }

	for result := range allResponse {
		count := 0
		total := 2
		direction := []int{0, 1}

		last := allResponse[result][0]
		next := allResponse[result][1]

		if last[0]-next[0] != direction[1] || last[1]-next[1] != direction[1] {
			count += 1
			direction[0], direction[1] = last[0]-next[0], last[1]-next[1]
		}

		for i := 2; i < len(allResponse[result]); i++ {
			last := allResponse[result][i-1]
			next := allResponse[result][i]

			if last[0]-next[0] != direction[0] || last[1]-next[1] != direction[1] {
				count += 1
				direction[0], direction[1] = last[0]-next[0], last[1]-next[1]
			}

			total += 1
		}

		// fmt.Println(total, count, len(allResponse[result]))
		response = min(response, total+count*1000)
	}
	// fmt.Println(response, "response")
	return response
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
