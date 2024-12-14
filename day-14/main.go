package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func stringToInt(char string) int {
	val, _ := strconv.Atoi(char)
	return val
}

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
func isInBound(matrix [][]string, row, col int) bool {
	return row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0])
}

func getAnswer(path string, times int) int {
	robots := parseStr(path)
	matrix := [][]int{}
	for row := 0; row < 103; row++ {
		matrix = append(matrix, []int{})
		for col := 0; col < 101; col++ {
			matrix[row] = append(matrix[row], 0)
		}
	}
	total := 0
	for _, robot := range robots {
		matrix[robot.y][robot.x] += 1
	}
	for _, robot := range robots {
		matrix[robot.y][robot.x] -= 1
		yres := int(math.Abs(float64((robot.vy * times) % len(matrix))))
		xres := int(math.Abs(float64((robot.vx * times) % len(matrix[0]))))
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
				robot.x += len(matrix[0]) - xres
			} else {
				robot.x -= xres
			}
		}
		if robot.vx > 0 {
			if robot.x+xres >= len(matrix[0]) {
				robot.x -= len(matrix[0]) - xres
			} else {
				robot.x += xres
			}
		}
		matrix[robot.y][robot.x] += 1
	}
	one, two, three, four := 0, 0, 0, 0
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] <= 0 {
				continue
			}
			if row < len(matrix)/2 && col < len(matrix[0])/2 {
				one += matrix[row][col]
			}
			if row < len(matrix)/2 && col > len(matrix[0])/2 {
				two += matrix[row][col]
			}
			if row > len(matrix)/2 && col < len(matrix[0])/2 {
				three += matrix[row][col]
			}
			if row > len(matrix)/2 && col > len(matrix[0])/2 {
				four += matrix[row][col]
			}
		}
	}
	// fmt.Println(one * two * three * four)
	fmt.Println(matrix)
	return total
}
func getStringFromNum(char string) string {
	num, _ := strconv.Atoi(char)
	return fmt.Sprintf("%d", num)
}
func main() {
	path := "sample"
	times := 100

	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	if len(os.Args) > 2 {
		times = stringToInt(os.Args[2])
	}

	result := getAnswer(path, times)

	fmt.Print(result)
}
