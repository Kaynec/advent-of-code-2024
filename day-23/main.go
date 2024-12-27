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

	locks := []string{}

	mapping := map[string]map[string]bool{}

	for _, num := range strings.Split(str, "\r\n") {
		splitted := strings.Split(num, "-")
		first, second := splitted[0], splitted[1]

		if mapping[first] != nil {

			mapping[first][second] = true
		} else {
			mapping[first] = map[string]bool{
				second: true,
			}
		}

		if mapping[second] != nil {

			mapping[second][first] = true
		} else {
			mapping[second] = map[string]bool{
				first: true,
			}
		}

	}

	threeLinesXoXo := []string{}
	for key := range mapping {

		for key2 := range mapping[key] {
			for key3 := range mapping[key2] {

				if key[0] != 't' && key2[0] != 't' && key3[0] != 't' || !mapping[key3][key] || key3 == key {
					continue
				}

				keySlice := strings.Split(fmt.Sprintf("%s%s%s", key, key2, key3), "")
				sort.Strings(keySlice)
				key := strings.Join(keySlice, "")
				idx := slices.Index(threeLinesXoXo, key)
				if idx == -1 {
					threeLinesXoXo = append(threeLinesXoXo, key)
				}
			}
		}

	}

	fmt.Println(len(threeLinesXoXo))
	return locks
}

func main() {
	parseStr(os.Args[1])
}
