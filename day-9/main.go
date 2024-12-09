package main

import (
	"fmt"
	"os"
	"slices"
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

// func compareCords(){

// }

func findAllSubstringNotDot(str []any) [][]int {

	res := [][]int{}

	for i := 0; i < len(str); i++ {
		if str[i] == "." {
			continue
		}
		for j := i; j < len(str); j++ {
			if str[j] != str[i] || j == len(str)-1 {
				if j == len(str)-1 {

					res = append(res, []int{i, len(str)})
					i = j
				} else {
					res = append(res, []int{i, j})
					i = j - 1
				}
				break
			}
		}
	}

	return res

}

func findAllSubstring(str []any) [][]int {

	res := [][]int{}

	for i := 0; i < len(str)-1; i++ {
		if str[i] == "." {

			for j := i; j < len(str)-1; j++ {
				if str[j] != "." {
					res = append(res, []int{i, j})
					i = j
					break
				}
			}

		}

	}

	return res

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
	eCord := findAllSubstring(str)
	cords := findAllSubstringNotDot(str)

	slices.Reverse(cords)

	for _, cord := range cords {

		cordItem := str[cord[0]]

		for index, e := range eCord {
			len := e[1] - e[0]
			diff := cord[1] - cord[0]
			if diff <= len && cord[0] > e[0] {
				for i := 0; i < diff; i++ {
					str[i+e[0]] = cordItem
					str[i+cord[0]] = "."
				}

				if len == diff {
					fmt.Println(len == diff)
					eCord = append(eCord[0:index], eCord[index+1:]...)
				} else {
					e[0] = e[0] + cord[1] - cord[0]
				}
				break

			}
		}
	}

	total := 0

	for multiplier, value := range str {
		if value == "." {
			continue
		}
		if s, ok := value.(int); ok {
			total += multiplier * s
		}
	}

	fmt.Println(total)
	return 0

}
func main() {
	path := "sample"
	if os.Args[1:] != nil && os.Args[1:][0] != "" {
		path = os.Args[1:][0]
	}
	fmt.Println(answer(path))
}
