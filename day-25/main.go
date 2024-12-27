package main

import (
	"fmt"
	"os"
	"strings"
)

func isLock(lockOrKey string) bool {
	for _, line := range strings.Split(lockOrKey, "\r\n")[:1] {
		for _, element := range strings.Split(line, "") {
			if element != "#" {
				return false
			}
		}
	}
	return true
}

func extractLockValues(lockOrKey string) []int {
	count := []int{0, 0, 0, 0, 0}

	for _, line := range strings.Split(lockOrKey, "\r\n")[1:] {

		for index, element := range strings.Split(line, "") {
			if element == "#" {
				count[index]++
			}
		}
	}
	return count
}
func extractKeyValues(lockOrKey string) []int {
	count := []int{-1, -1, -1, -1, -1}

	lockOrKeyLines := strings.Split(lockOrKey, "\r\n")

	for i := len(lockOrKeyLines) - 1; i >= 0; i-- {
		line := lockOrKeyLines[i]
		for index, element := range strings.Split(line, "") {
			if element == "#" {
				count[index]++
			}
		}
	}
	return count
}

func parseStr(path string) ([][]int, [][]int) {
	data, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := strings.TrimSpace(string(data))

	locks := [][]int{}
	keys := [][]int{}

	for _, lockOrKey := range strings.Split(str, "\r\n\r\n") {

		if isLock(lockOrKey) {
			locks = append(locks, extractLockValues(lockOrKey))
		} else {
			keys = append(keys, extractKeyValues(lockOrKey))
		}

	}

	return locks, keys
}

const LockLength = 5

func main() {
	locks, keys := parseStr(os.Args[1])

	count := len(locks) * len(keys)
	for _, locklist := range locks {

	currKeyLoop:
		for _, keyList := range keys {

			for i := 0; i < LockLength; i++ {
				if keyList[i]+locklist[i] > 5 {
					count--
					continue currKeyLoop
				}
			}
		}
	}
	fmt.Println(count)
}
