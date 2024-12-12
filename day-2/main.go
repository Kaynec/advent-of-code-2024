package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func mapLinesToNums(path string) [][]int {

	val, err := os.ReadFile(fmt.Sprintf("%s.txt", path))

	if err != nil {
		fmt.Println("Error : Something Wrong Happend Buddy")
	}

	lines := strings.Split(strings.TrimSpace(string(val)), "\r\n")

	newLines := make([][]int, len(lines))

	for lineId, line := range lines {

		for _, field := range strings.Fields(line) {

			val, _ := strconv.Atoi(field)

			newLines[lineId] = append(newLines[lineId], val)
		}
	}
	return newLines
}

func partOneAnswer() {
	newLines := mapLinesToNums("sample")
	sum := 0

	for _, line := range newLines {
		safe := true

		direction := "increase"
		if line[0] < line[1] {
			direction = "decrease"
		}
		for i := 0; i < len(line)-1; i++ {
			diff := int(math.Abs(float64(line[i] - line[i+1])))
			if diff > 3 || diff < 1 {
				safe = false
			}

			dir := "increase"

			if line[i] < line[i+1] {
				dir = "decrease"
			}
			if direction != dir {
				safe = false
			}

		}
		if safe {
			sum++
		}
	}
	fmt.Println(sum)
}

func returnDirection(arr []int, i int) string {

	if arr[i] < arr[i+1] {
		return "decrease"
	}
	return "increase"

}

func isSafe(line []int) bool {
	badLevelCount := 0

	direction := returnDirection(line, 0)
	for i := 0; i < len(line)-1; i++ {
		if line[i] == line[i+1] {
			continue
		}

		dir := returnDirection(line, i)

		if direction != dir {
			badLevelCount++
		}

		diff := int(math.Abs(float64(line[i] - line[i+1])))

		if diff > 3 {
			return false
		}

	}
	return badLevelCount <= 1

}

func partTwoAnswer() {
	// newLines, sum := mapLinesToNums("sample"), 0
	newLines, sum := mapLinesToNums("input"), 0

	for _, line := range newLines {
		if isSafe(line) {
			sum++
		}
	}
	fmt.Println(sum)
}

func main() {
	partTwoAnswer()
}
