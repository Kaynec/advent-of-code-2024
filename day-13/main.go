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
			X: stringToInt(strings.Split(strings.Split(third, "Prize: X=")[1], ",")[0]),
			Y: stringToInt(strings.Split(strings.TrimSpace(strings.Split(third, "Prize: X=")[1]), "Y=")[1]),
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
		buttonAndTarget := matrix[i]
		smallOffset := 0

		start := buttonAndTarget.Target.X / max(buttonAndTarget.Button1.X, buttonAndTarget.Button2.X)
		for i := start; i > 0; i-- {
			res := buttonAndTarget.Target.X - i*max(buttonAndTarget.Button1.X, buttonAndTarget.Button2.X)

			if res%min(buttonAndTarget.Button1.X, buttonAndTarget.Button2.X) == 0 {
				smallOffset = res / min(buttonAndTarget.Button1.X, buttonAndTarget.Button2.X)

				if i*max(buttonAndTarget.Button1.X, buttonAndTarget.Button2.X)+smallOffset*min(buttonAndTarget.Button1.X, buttonAndTarget.Button2.X) != buttonAndTarget.X {
					continue
				}
				if buttonAndTarget.Button1.X > buttonAndTarget.Button2.X {
					if i*buttonAndTarget.Button1.Y+smallOffset*buttonAndTarget.Button2.Y != buttonAndTarget.Y {
						continue
					}
					total += i*3 + smallOffset
				} else {
					if i*buttonAndTarget.Button2.Y+smallOffset*buttonAndTarget.Button1.Y != buttonAndTarget.Y {
						continue
					}
					total += i + smallOffset*3

				}
				break
			}

		}
	}
	fmt.Println(total)
	// total := 0

	return 0
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
