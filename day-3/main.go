package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseStr(path string, regexStr string) []string {
	val, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := string(val)

	regex := regexp.MustCompile(regexStr)

	mulSlice := regex.FindAllString(str, -1)

	return mulSlice
}

func returnMul(slice string) int {
	slice = strings.TrimSpace(slice)

	firstSplit := strings.Split(slice, "mul(")

	numsAndEndParantheses := strings.Split(firstSplit[1], ")")
	nums := numsAndEndParantheses[0]

	firstNumStr, secondNumStr := strings.Split(nums, ",")[0], strings.Split(nums, ",")[1]

	firstNum, _ := strconv.Atoi(firstNumStr)

	secondNum, _ := strconv.Atoi(secondNumStr)

	return (firstNum * secondNum)
}

func partOneAnswer(path string) {
	mulSlice := parseStr(path, `mul\(\s*-?\d+\s*,\s*-?\d+\s*\)`)

	sum := 0

	for _, slice := range mulSlice {

		sum = sum + returnMul(slice)
	}

}

func partTwoAnswer(path string) int {
	mulSlice := parseStr(path, `(mul\(\s*-?\d+\s*,\s*-?\d+\s*\)|don't\(\)|do\(\))`)

	shouldMul := true

	sum := 0

	for _, value := range mulSlice {

		if value == "do()" {
			shouldMul = true
			continue
		}
		if value == "don't()" {
			shouldMul = false
			continue
		}

		if shouldMul {
			sum += returnMul(value)
		}

	}

	return sum
}

func main() {
	partTwoAnswer("input")
}
