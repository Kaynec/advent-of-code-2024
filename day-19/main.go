package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func isPossible(assignement []string, combination string) bool {

	newString := strings.TrimSpace(combination)

	for range combination {

		for _, assign := range assignement {
			assign = strings.TrimSpace(assign)
			newString = strings.Replace(newString, assign, "", 1)
		}
	}

	return len(newString) == 0
}

func parseStr(path string) [][]string {
	data, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := strings.TrimSpace(string(data))
	matrix := [][]string{}
	assignement := strings.Split(strings.TrimSpace(strings.Split(str, "\r\n\r\n")[0]), ",")

	sort.Slice(assignement, func(i, j int) bool {
		return len(assignement[i]) > len(assignement[j])
	})
	secondPart := strings.Split(str, "\r\n\r\n")[1]
	sum := 0
	for _, line := range strings.Split(secondPart, "\r\n") {
		if ok := isPossible(assignement, line); ok {
			sum += 1
		}
	}
	fmt.Println(sum)
	return matrix
}
func main() {
	parseStr(os.Args[1])
}
