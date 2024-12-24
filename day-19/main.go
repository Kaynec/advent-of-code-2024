package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

)

func countValidCases(assignement []string, combination string, cache map[string]int) (count int) {
	if val, ok := cache[combination]; ok {
		return val
	}
	// Base case: if combination is empty, return true
	if len(combination) == 0 {
		// fmt.Println("Found")
		cache[combination] = max(count, 1)
		return max(count, 1)
	}

	// Iterate through each available pattern
	for _, assign := range assignement {
		// Check if the pattern matches at any position in combination
		if strings.HasPrefix(combination, assign) {
			// Recursively check the remaining part of the combination
			key := fmt.Sprintf("%s,%s", combination, assign)

			num, ok := cache[key]

			if ok {
				count += num
			} else {
				cache[key] = countValidCases(assignement, combination[len(assign):], cache)
				count += cache[key]
			}

		}
	}
	// If no patterns matched, return false
	return count
}

func parseStr(path string) [][]string {
	data, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := strings.TrimSpace(string(data))
	matrix := [][]string{}
	assignement := strings.Split(strings.TrimSpace(strings.Split(str, "\r\n\r\n")[0]), ",")
	for i := range assignement {
		assignement[i] = strings.TrimSpace(assignement[i])
	}

	sort.Slice(assignement, func(i, j int) bool {
		return len(assignement[i]) > len(assignement[j])
	})
	secondPart := strings.Split(str, "\r\n\r\n")[1]
	sum := 0
	cache := map[string]int{}
	for _, line := range strings.Split(secondPart, "\r\n") {
		sum += countValidCases(assignement, line, cache)
	}
	fmt.Println(sum)
	return matrix
}
func main() {
	parseStr(os.Args[1])
}
