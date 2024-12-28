package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"

)

func parseStr(path string) []string {
	data, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := strings.TrimSpace(string(data))

	network := map[string][]string{}

	for _, num := range strings.Split(str, "\r\n") {
		splitted := strings.Split(num, "-")
		first, second := strings.TrimSpace(splitted[0]), strings.TrimSpace(splitted[1])

		if len(network[first]) <= 0 {
			network[first] = []string{second}
		}
		if len(network[second]) <= 0 {
			network[second] = []string{first}
		}

		network[second] = append(network[second], first)
		network[first] = append(network[first], second)

	}

	threeLinesXoXo := map[string]bool{}
	for x := range network {
		if !strings.HasPrefix(x, "t") {
			continue
		}
		for _, y := range network[x] {

			for _, z := range network[y] {

				if !slices.Contains(network[z], x) || z == x {
					continue
				}

				keys := []string{x, y, z}

				sort.Strings(keys)

				keySlice := strings.Split(fmt.Sprintf("%s", keys), "")

				key := strings.Join(keySlice, "")

				threeLinesXoXo[key] = true
			}
		}

	}

	fmt.Println(len(threeLinesXoXo))
	return strings.Split(str, "")
}

func main() {
	parseStr(os.Args[1])
}
