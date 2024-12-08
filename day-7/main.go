package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseStr(path string) []string {
	val, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := string(val)

	strSlice := strings.Split(str, "\r\n")

	return strSlice
}
func glueNumbers(a int, b int) int {
	strRepresented := strconv.Itoa(a) + strconv.Itoa(b)
	val, _ := strconv.Atoi(strRepresented)
	return val
}
func walk(numbers []int, index int, target, total int) int {
	if index >= len(numbers) {
		return total
	}

	firstPath := walk(numbers, index+1, target, total+numbers[index])
	secondPath := walk(numbers, index+1, target, total*numbers[index])
	thridPath := walk(numbers, index+1, target, glueNumbers(total, numbers[index]))

	if firstPath == target || secondPath == target || total == target || thridPath == target {
		return target
	}
	return total
}
func getCorrectEquationsCount(numbers []int, target int) bool {
	res := walk(numbers, 1, target, numbers[0])
	return res == target
}
func partOne(path string) {
	roots := parseStr(path)

	total := 0
	for _, root := range roots {
		fields := strings.Split(root, ":")
		t, _ := strconv.Atoi(fields[0])
		numbers := make([]int, 0)
		for _, num := range strings.Fields(fields[1]) {
			n, _ := strconv.Atoi(num)
			numbers = append(numbers, n)
		}
		if getCorrectEquationsCount(numbers, t) {
			total += t
		}
	}

	fmt.Println(total)
}
func main() {
	partOne("sample")
}
