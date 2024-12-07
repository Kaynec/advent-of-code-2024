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

	mulSlice := strings.Split(str, "\r\n")

	return mulSlice
}

// slice =  11 6 16 20
// slice =  81 40 27
func walk(numbers []int, index int, res *[][]string, operands []string) {
	if index == len(numbers)-1 {
		*res = append(*res, operands)
		return
	}
	walk(numbers, index+1, res, append(operands, "+"))
	walk(numbers, index+1, res, append(operands, "*"))
}
func getCorrectEquationsCount(numbers []int, target int) bool {
	operandsMatrix := [][]string{}

	walk(numbers, 0, &operandsMatrix, []string{})

	for _, operands := range operandsMatrix {
		sum := numbers[0]
		numIdx := 1
		for _, operand := range operands {
			if operand == "+" {
				sum += numbers[numIdx]
				numIdx++
			}
			if operand == "*" {
				sum *= numbers[numIdx]
				numIdx++
			}
		}
		if sum == target {
			return true
		}
	}
	return false
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
	partOne("input")
}
