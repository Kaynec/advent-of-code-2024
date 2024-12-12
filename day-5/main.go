package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseStr(path string) []string {
	val, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := string(val)

	mulSlice := strings.Split(str, "\r\n\r\n")

	return mulSlice
}

func isGoodOrder(order []string, endMap map[string][]string) bool {

	before := make(map[string]bool)
	existsInOrder := make(map[string]bool)

	for _, num := range order {

		existsInOrder[num] = true

	}

	for _, num := range order {

		innerNums, ok := endMap[num]

		if ok {
			for _, innerNum := range innerNums {
				if !before[innerNum] && existsInOrder[innerNum] {
					return false
				}
			}
		}

		before[num] = true

	}
	return true
}
func fixOrder(order []string, table map[string][]string) []string {
	before := make(map[string]bool)

	for i := 0; i < len(order); i++ {

		num := order[i]

		indecis := []int{}
		for _, innerNum := range table[num] {
			if before[innerNum] {
				innerNumIdx := slices.Index(order, innerNum)
				indecis = append(indecis, innerNumIdx)
			}

		}
		if len(indecis) > 0 {
			fmt.Println(i, indecis)
			slices.Sort(indecis)
			indeci := indecis[0]
			idx := slices.Index(order, num)
			order[indeci], order[idx] = order[idx], order[indeci]
			i--
		}
		before[num] = true
	}
	fmt.Println(order)
	return order
}

func getAnswer(path string) {

	parts := parseStr(path)
	first, second := parts[0], parts[1]

	startMap := make(map[string][]string)
	endMap := make(map[string][]string)

	for _, line := range strings.Split(first, "\r\n") {
		nums := strings.Split(line, "|")

		_, ok := startMap[nums[0]]
		if ok {

			startMap[nums[0]] = append(startMap[nums[0]], nums[1])
		} else {
			startMap[nums[0]] = []string{nums[1]}
		}

		_, ok2 := endMap[nums[1]]
		if ok2 {

			endMap[nums[1]] = append(endMap[nums[1]], nums[0])
		} else {
			endMap[nums[1]] = []string{nums[0]}
		}
	}
	strings.Split(second, "\r\n")
	orders := strings.Split(second, "\r\n")

	sum := 0
	for _, order := range orders {

		orderParsed := strings.Split(order, ",")

		if isGoodOrder(orderParsed, endMap) {
			continue
		} else {
			newOrders := fixOrder(orderParsed, startMap)
			val, _ := strconv.Atoi(newOrders[len(newOrders)/2])
			sum += val

		}
	}
	fmt.Println(sum)
}

func main() {
	getAnswer("sample")
}
