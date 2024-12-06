package main

import (
	"fmt"
	"os"
	"strings"
)

func checkIsBound(slice [][]string, nums []int) bool {

	fmt.Println(nums)

	return false
}

func checkXmas(slice [][]string, row int, col int) {

	numsSlice := make([][]int, 0)

	numsSlice = append(numsSlice, []int{row, col}, []int{row, col + 1})
	fmt.Println(numsSlice)

	// checkIsBound(slice, &{[0,0]} )

	// SAMX
	// if strings.Join(slice[row], "")[col-4:col] == "SAMX" {
	// 	fmt.Println("1")
	// }
	// // XMAS
	// if strings.Join(slice[row], "")[col:col+4] == "XMAS" {
	// 	fmt.Println("1")
	// }
	// if row+3 > len(slice) {
	// 	return
	// }
	// fmt.Println(
	// 	checkIsBound(slice, row, col+3),
	// checkIsBound(slice, row, col-3),
	// checkAngle(slice, row+3, col),
	// checkAngle(slice, row-3, col),

	// checkAngle(slice, row-3, col-3),
	// checkAngle(slice, row-3, col-3),

	// checkAngle(slice, row+3, col+3),
	// checkAngle(slice, row+3, col+3),
}

func parseStr(path string) [][]string {
	val, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := string(val)

	var sliced [][]string

	mulSlice := strings.Split(str, "\r\n")

	for _, rowVal := range mulSlice {

		sliced = append(sliced, strings.Split(rowVal, ""))
	}

	return sliced
}

func partOneAnswer(path string) {
	mulSlice := parseStr(path)

	sum := 0

	for row := range mulSlice {

		for col := range mulSlice[row] {

			if mulSlice[row][col] == "X" {
				checkXmas(mulSlice, row, col)
			}

		}
	}

	fmt.Println(sum)

}

// func partTwoAnswer(path string) int {
// 	mulSlice := parseStr(path, `(mul\(\s*-?\d+\s*,\s*-?\d+\s*\)|don't\(\)|do\(\))`)

// 	shouldMul := true

// 	sum := 0

// 	for _, value := range mulSlice {

// 		if value == "do()" {
// 			shouldMul = true
// 			continue
// 		}
// 		if value == "don't()" {
// 			shouldMul = false
// 			continue
// 		}

// 		if shouldMul {
// 			sum += returnMul(value)
// 		}

// 	}

// 	return sum
// }

func main() {
	partOneAnswer("sample")
}
