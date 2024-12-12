package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MULTIPLIER = 2024

type Stone struct {
	count int
}

func stringToInt(char string) int {
	val, _ := strconv.Atoi(char)
	return val
}
func parseStr(path string) []string {
	data, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := strings.TrimSpace(string(data))
	return strings.Fields(str)
}
func createOrUpdateStone(table map[string]Stone, key string, newStoneCount int) {
	if val, exists := table[key]; exists {
		val.count += newStoneCount
		table[key] = val
	} else {
		table[key] = Stone{count: newStoneCount}
	}
}
func getAnswer(path string, times int) int {
	stones := parseStr(path)

	total := 0
	table := make(map[string]Stone)

	for _, stone := range stones {
		createOrUpdateStone(table, stone, 1)
	}

	for i := 0; i < times; i++ {
		tempTable := make(map[string]Stone)

		for key, value := range table {
			switch {
			case key == "0":
				createOrUpdateStone(tempTable, "1", value.count)
			case len(key)%2 == 0:
				leftAndRight := strings.Split(key, "")
				mid := len(leftAndRight) / 2
				left := getStringFromNum(strings.Join(leftAndRight[:mid], ""))
				right := getStringFromNum(strings.Join(leftAndRight[mid:], ""))
				createOrUpdateStone(tempTable, left, value.count)
				createOrUpdateStone(tempTable, right, value.count)
			default:
				newKey := fmt.Sprintf("%d", stringToInt(key)*MULTIPLIER)
				createOrUpdateStone(tempTable, newKey, value.count)
			}
		}

		table = tempTable
	}

	for _, val := range table {
		total += val.count
	}
	return total
}
func getStringFromNum(char string) string {
	num, _ := strconv.Atoi(char)
	return fmt.Sprintf("%d", num)
}
func main() {
	path := "sample"
	times := 25

	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	if len(os.Args) > 2 {
		times = stringToInt(os.Args[2])
	}

	result := getAnswer(path, times)

	fmt.Print(result)
}
