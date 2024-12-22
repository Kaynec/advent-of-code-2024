package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

)

func stringToInt(char string) int {
	val, _ := strconv.Atoi(char)
	return val
}

type Target struct {
	X int
	Y int
}
type Button struct {
	X int
	Y int
}

type ButtonAndTarget struct {
	Button1 Button
	Button2 Button
	Target
}

const ADDAMOUNT = 10000000000000

func getNumsFromStr(str string, letter string) (int, int) {
	nums := strings.TrimSpace(strings.Split(str, fmt.Sprintf("Button %s: X+", letter))[1])
	a1 := strings.Split(nums, ",")[0]
	a2 := strings.Split(strings.TrimSpace(strings.Split(nums, ",")[1]), "Y+")[1]

	return stringToInt(a1), stringToInt(a2)
}

func parseStr(path string) []ButtonAndTarget {
	data, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := strings.TrimSpace(string(data))
	matrix := []ButtonAndTarget{}

	for _, line := range strings.Split(str, "\r\n\r\n") {

		innerLines := strings.Split(line, "\r\n")
		first, second, third := innerLines[0], innerLines[1], innerLines[2]

		x, y := getNumsFromStr(first, "A")
		x2, y2 := getNumsFromStr(second, "B")
		Button1 := Button{
			X: x,
			Y: y,
		}
		Button2 := Button{
			X: x2,
			Y: y2,
		}
		target := Target{
			X: stringToInt(strings.Split(strings.Split(third, "Prize: X=")[1], ",")[0]) + +ADDAMOUNT,
			Y: stringToInt(strings.Split(strings.TrimSpace(strings.Split(third, "Prize: X=")[1]), "Y=")[1]) + +ADDAMOUNT,
		}
		matrix = append(matrix, ButtonAndTarget{
			Target:  target,
			Button1: Button1,
			Button2: Button2,
		})
	}
	return matrix
}
func getAnswer(path string) int {
	matrix := parseStr(path)
	total := 0
	for i := 0; i < len(matrix); i++ {
		// using cramer's rule just black magic
		buttonAndTarget := matrix[i]
		D := (buttonAndTarget.Button1.X * buttonAndTarget.Button2.Y) - (buttonAndTarget.Button1.Y * buttonAndTarget.Button2.X)
		X := (buttonAndTarget.Target.X * buttonAndTarget.Button2.Y) - (buttonAndTarget.Target.Y * buttonAndTarget.Button2.X)
		Y := (buttonAndTarget.Button1.X * buttonAndTarget.Target.Y) - (buttonAndTarget.Target.X * buttonAndTarget.Button1.Y)

		X = X / D
		Y = Y / D
		if (X*buttonAndTarget.Button1.X)+(Y*buttonAndTarget.Button2.X) == buttonAndTarget.Target.X && (X*buttonAndTarget.Button1.Y)+Y*(buttonAndTarget.Button2.Y) == buttonAndTarget.Target.Y {
			total += X*3 + Y
		}
	}
	return total
}
func main() {
	path := "sample"

	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	result := getAnswer(path)

	fmt.Print(result)
}
