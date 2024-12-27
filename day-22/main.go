package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseStr(path string) []int {
	data, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := strings.TrimSpace(string(data))

	locks := []int{}

	for _, num := range strings.Split(str, "\r\n") {
		val, err := strconv.Atoi(strings.TrimSpace(num))
		if err == nil {
			locks = append(locks, val)
		}
	}

	return locks
}

const LockLength = 5

type MonkeyNum struct {
	num     int
	diff    []int
	last    int
	hashMap map[string]bool
}

const Prune = 16777216

func getLastDigitOfNum(num int) int {
	str := strconv.Itoa(num)

	newNum, _ := strconv.Atoi(string(str[len(str)-1]))

	return newNum
}

func main() {
	rawnumbers := parseStr(os.Args[1])
	numbers := []MonkeyNum{}

	hashMap := map[string]int{}

	for _, num := range rawnumbers {
		numbers = append(numbers, MonkeyNum{
			num:  num,
			diff: []int{0, 0, 0, 0},
		})
	}

	for i := 1; i <= 2000; i++ {
		for j, num := range numbers {
			rawN := num.num
			secret := rawN
			newSecret := (((rawN * 64) ^ secret) % Prune)
			newSecret = int(math.Round(float64(newSecret/32))) ^ newSecret
			newSecret = (((newSecret % Prune) * 2048) ^ newSecret) % Prune
			numbers[j].num = newSecret
			rawN = newSecret
			tmpLast := num.last
			numbers[j].last = getLastDigitOfNum(rawN)

			currDiff := 0

			if i == 1 {
				currDiff = numbers[j].last - getLastDigitOfNum(num.num)
				numbers[j].hashMap = map[string]bool{}

			} else {
				currDiff = numbers[j].last - tmpLast
			}

			numbers[j].diff = append(numbers[j].diff, currDiff)
			numbers[j].diff = numbers[j].diff[1:]

			if i < 4 {
				continue
			}

			diffList := numbers[j].diff
			id := len(diffList) - 1

			key := fmt.Sprintf("%d,%d,%d,%d", diffList[id-3], diffList[id-2], diffList[id-1], diffList[id])

			_, isFirstTime := numbers[j].hashMap[key]

			if !isFirstTime {
				numbers[j].hashMap[key] = true

				hashMap[key] += numbers[j].last
			}

		}
	}

	total := 0

	for _, value := range hashMap {
		if value > total {
			total = value
		}
	}

	fmt.Println(total)
}
