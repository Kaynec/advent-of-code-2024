package main

import (
	"fmt"
	"math"
	"os"
	"sort"

	// "sort"
	"strconv"
	"strings"
)

func stringToInt(char string) int {
	val, _ := strconv.Atoi(char)
	return val
}

const ROW = 103
const COL = 101

type Robot struct {
	x  int
	y  int
	vx int
	vy int
}

func parseStr(path string) []Robot {
	data, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := strings.TrimSpace(string(data))

	matrix := []Robot{}

	for _, line := range strings.Split(str, "\r\n") {
		fields := strings.Fields(line)
		first, second := fields[0], fields[1]
		axis := strings.Split(strings.Split(first, "=")[1], ",")
		velocity := strings.Split(strings.Split(second, "=")[1], ",")
		matrix = append(matrix, Robot{
			x:  stringToInt(axis[0]),
			y:  stringToInt(axis[1]),
			vx: stringToInt(velocity[0]),
			vy: stringToInt(velocity[1]),
		})
	}
	return matrix
}

func isInBound[S ~[][]E, E comparable](matrix S, row, col int) bool {
	return row >= 0 && row < ROW && col >= 0 && col < COL
}

func getShape(matrix [][]int, row, col int, res *[][]int, table map[string]bool) {
	key := fmt.Sprintf("%d,%d", row, col)

	if !isInBound(matrix, row, col) {
		return
	}
	if table[key] || matrix[row][col] <= 0 {
		return
	}
	*res = append(*res, []int{row, col})
	table[key] = true

	getShape(matrix, row+1, col, res, table)
	getShape(matrix, row-1, col, res, table)
	getShape(matrix, row, col+1, res, table)
	getShape(matrix, row, col-1, res, table)

}
func returnGetShapeResult(matrix [][]int, row, col int, table map[string]bool) (res [][]int) {
	getShape(matrix, row, col, &res, table)
	return res

}

const MAX_TIMES = 11000

func getAnswer(path string) int {
	robots := parseStr(path)

	total := 0

	for times := 10; times < 11000; times++ {
		res := [][][]int{}
		table := make(map[string]bool)
		matrix := [][]int{}
		for row := 0; row < ROW; row++ {
			matrix = append(matrix, []int{})
			for col := 0; col < COL; col++ {
				matrix[row] = append(matrix[row], 0)
			}
		}

		for _, robot := range robots {
			yres := int(math.Abs(float64((robot.vy * times) % len(matrix))))
			xres := int(math.Abs(float64((robot.vx * times) % COL)))
			if robot.vy < 0 {
				if robot.y-yres < 0 {
					robot.y += len(matrix) - yres
				} else {
					robot.y -= yres
				}
			}
			if robot.vy > 0 {
				if robot.y+yres >= len(matrix) {
					robot.y -= len(matrix) - yres
				} else {
					robot.y += yres
				}
			}
			if robot.vx < 0 {
				if robot.x-xres < 0 {
					robot.x += COL - xres
				} else {
					robot.x -= xres
				}
			}
			if robot.vx > 0 {
				if robot.x+xres >= COL {
					robot.x -= COL - xres
				} else {
					robot.x += xres
				}
			}
			matrix[robot.y][robot.x] += 1
		}
		for row := range matrix {
			res = append(res, [][]int{})
			for col := range matrix[row] {
				key := fmt.Sprintf("%d,%d", row, col)

				if table[key] {
					continue
				}

				result := returnGetShapeResult(matrix, row, col, table)
				if len(result) > 1 {
					res = append(res, result)

				}
			}
		}

	outer:
		for row := range res {
			sort.Slice(res[row], func(i, j int) bool {
				if res[row][i][0] == res[row][j][0] {
					if res[row][i][1] > res[row][j][1] {
						return false
					} else {
						return true
					}
				} else {
					return res[row][i][0] < res[row][j][0]
				}
			})

			if len(res[row]) < 12 {
				continue
			}

			lastValue := []int{}
			firstValue := []int{}
			str := []string{}
			cords := [][][]int{}
			maxDistance := 0

			if times != 7083 {
				continue
			}

			for index, matrix := range res[row] {

				if index == 0 {
					firstValue = []int{matrix[0], matrix[1]}
					lastValue = []int{matrix[0], matrix[1]}
					str = append(str, "1")
					cords = append(cords, [][]int{})
					cords[len(cords)-1] = append(cords[len(cords)-1])
					cords[len(cords)-1] = append(cords[len(cords)-1], []int{matrix[0], matrix[1]})
					continue
				}

				if matrix[0] > lastValue[0] {
					str = append(str, "\r\n")
					firstValue = []int{matrix[0], matrix[1]}
					cords = append(cords, [][]int{})
				}
				maxDistance = max(int(math.Abs(float64(lastValue[1]-firstValue[1]))), maxDistance)
				str = append(str, "1")
				lastValue = []int{matrix[0], matrix[1]}
				cords[len(cords)-1] = append(cords[len(cords)-1], []int{matrix[0], matrix[1]})
			}

			for index, cord := range cords {
				padding := 0
				sort.Slice(cord, func(i, j int) bool { return i < j })
				firstItem, lastItem := cord[0][1], cord[len(cord)-1][1]

				if index == 0 {
					continue
				}
				lastListFirstItem, lastListLastItem := cords[index-1][0], cords[index-1][len(cords[index-1])-1]

				// if length is the same for more than 3 times, then it is a line

				// it's bigger
				if firstItem < lastListFirstItem[1] {
					padding = (lastItem - firstItem) - len(cords[index-1])

					if lastListFirstItem[1] != firstItem+padding {
						break outer
					}
					if lastListLastItem[1] != lastItem-padding {
						break outer
					}

				}
				// it's smaller
				if firstItem > lastListFirstItem[1] {
					padding = (len(cords[index-1]) - (lastItem - firstItem)) / 2

					if firstItem != lastListFirstItem[1]+padding {
						break outer

					}
					if lastItem != lastListLastItem[1]-padding {
						break outer
					}
				}
			}
			for index, options := range res[row][:len(res[row])-2] {
				if res[row][index+1][1]-options[1] > 1 {
					continue outer
				}
			}
			return times
		}

	}
	return total
}
func getStringFromNum(char string) string {
	num, _ := strconv.Atoi(char)
	return fmt.Sprintf("%d", num)
}
func main() {
	path := "sample"

	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	result := getAnswer(path)

	fmt.Print(result)
}
