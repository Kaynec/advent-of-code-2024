package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toNum(char string) int {
	val, _ := strconv.Atoi(char)
	return val
}
func toStr(val int) string {
	return fmt.Sprintf("%d", val)
}

func unshift[T any](slice *[]T) T {

	unshift := (*slice)[0]

	*slice = (*slice)[1:]

	return unshift
}

func parseStr(path string) [][]int {
	val, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := strings.Split(string(val), "")

	var lineMatrix [][]int

	for i := 0; i < len(str); i += 2 {
		if i+1 == len(str) {
			lineMatrix = append(lineMatrix, []int{toNum(str[i]), 0})
			break
		}

		lineMatrix = append(lineMatrix, []int{toNum(str[i]), toNum(str[i+1])})

	}

	return lineMatrix
}

func answer(path string) int {
	matrix := parseStr(path)

	str := []any{}

	emptyCord := []int{}
	for id, slice := range matrix {

		for i := 0; i < slice[0]; i++ {

			str = append(str, id)
		}
		len := len(str)
		for i := 0; i < slice[1]; i++ {

			str = append(str, ".")

			emptyCord = append(emptyCord, len+i)
		}
	}
	length := len(emptyCord)

	right := len(str) - 1
	for right >= 0 {
		if len(emptyCord) <= 0 {
			break
		}
		if str[right] == "." {
			emptyCord = append(emptyCord, right)
			right--
			continue
		}

		unshifted := unshift(&emptyCord)
		str[unshifted] = str[right]
		str[right] = str[unshifted]
		right--

	}

	//
	total := 0
	for multiplier, number := range str[:len(str)-length] {
		if s, ok := number.(int); ok {
			total += multiplier * s
		}
	}

	fmt.Println(total)

	return 0

}
func main() {
	answer("input")
}
