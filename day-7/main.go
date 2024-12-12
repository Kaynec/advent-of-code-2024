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
func walk(numbers []int, index int, target, total int) bool {
	if index > len(numbers)-1 {
		return total == target
	}

	firstPath := walk(numbers, index+1, target, total+numbers[index])
	secondPath := walk(numbers, index+1, target, total*numbers[index])

	thridPath := walk(numbers, index+1, target, glueNumbers(total, numbers[index]))
	return firstPath || secondPath || thridPath
}
func getCorrectEquationsCount(numbers []int, target int) bool {
	res := walk(numbers, 1, target, numbers[0])
	fmt.Println(res)
	return res
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

	path := "sample"
	if os.Args[1:] != nil && os.Args[1:][0] != "" {
		path = os.Args[1:][0]
	}
	partOne(path)
}
