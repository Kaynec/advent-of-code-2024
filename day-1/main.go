package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func mapLinesToNums(path string) ([]int, []int) {
	val, err := os.ReadFile(fmt.Sprintf("%s.txt", path))

	if err != nil {
		fmt.Println(err)
	}

	strValue := string(val)
	strValue = strings.TrimSpace(strValue)

	var firstSlice []int
	var secondSlice []int

	for _, line := range strings.Split(strValue, "\r\n") {
		line = strings.TrimSpace((line))
		newLine := strings.Split(line, "   ")

		first, second := newLine[0], newLine[1]

		firstNum, _ := strconv.Atoi(first)
		secondNum, _ := strconv.Atoi(second)

		firstSlice = append(firstSlice, firstNum)

		secondSlice = append(secondSlice, secondNum)

	}

	return firstSlice, secondSlice
}

func main() {

	firstSlice, secondSlice := mapLinesToNums("sample")

	sum := 0

	slices.Sort(firstSlice)
	slices.Sort(secondSlice)

	for i := 0; i < len(firstSlice); i++ {

		sum += int(math.Abs(float64(firstSlice[i] - secondSlice[i])))
	}
	fmt.Println(sum)
	fmt.Println(firstSlice, secondSlice)
	// fmt.Println(
	// 	mapLinesToNums("sample"),
	// )
	// splitFileByTwo("input")
}
