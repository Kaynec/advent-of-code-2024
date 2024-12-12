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

func parseStr(path string) []int {
	val, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := strings.TrimSpace(string(val))

	ints := []int{}
	for _, integ := range strings.Split(str, ",") {
		ints = append(ints, getNum(integ))
	}
	return ints
}

func getNum(char string) int {
	s, _ := strconv.Atoi(char)
	return s
}

func getAnswer(path string, times int) int {
	stones := parseStr(path)

	total := 0

	table := make(map[int]int)
	for i := 0; i < 9; i++ {
		table[i] = 0
	}
	for _, stone := range stones {
		table[stone] += 1
	}

	for i := 0; i < times; i++ {

		tmp := table[0]
		table[0] = table[1]

		for i := 1; i < 8; i++ {
			table[i] = table[i+1]
			if i != 6 {
				continue
			}
			table[6] = table[6] + tmp
		}
		table[8] = tmp
	}
	for _, value := range table {
		total += value
	}

	return total
}

func main() {
	path := "sample"
	times := 25
	if len(os.Args) > 1 {
		path = os.Args[1]
		if len(os.Args) > 2 {
			num, _ := strconv.Atoi(os.Args[2])
			times = num
		}
	}
	result := getAnswer(path, times)
	fmt.Println(result)
}
