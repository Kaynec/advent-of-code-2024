package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strings"
)

type Network = map[string]map[string]bool

// res2 is only for three connected computers xo xo
func recursoveFn(network Network, key string, cache map[string]bool, res *[]string) {
	if cache[key] || key == "" || len(network[key]) == 0 {
		return
	}

	cache[key] = true

	if !slices.Contains(*res, key) {
		*res = append(*res, key)
	}
	for newNetwork := range network[key] {
		for _, oldKey := range *res {
			if !network[oldKey][newNetwork] {
				return
			}
		}
		recursoveFn(network, newNetwork, cache, res)
	}

}

func parseStr(path string) Network {
	data, _ := os.ReadFile(fmt.Sprintf("%s.txt", path))
	str := strings.TrimSpace(string(data))

	network := Network{}

	for _, num := range strings.Split(str, "\r\n") {
		splitted := strings.Split(num, "-")
		first, second := strings.TrimSpace(splitted[0]), strings.TrimSpace(splitted[1])

		if len(network[first]) <= 0 {
			network[first] = map[string]bool{second: true}
		}
		if len(network[second]) <= 0 {
			network[second] = map[string]bool{first: true}
		}

		network[second][first] = true
		network[first][second] = true

	}
	return network
}

func part1(network Network) int {
	threeLinesXoXo := map[string]bool{}
	for x := range network {
		if !strings.HasPrefix(x, "t") {
			continue
		}
		for y := range network[x] {

			for z := range network[y] {

				if !network[z][x] {
					continue
				}

				keys := []string{x, y, z}

				sort.Strings(keys)

				key := strings.Join(strings.Split(fmt.Sprintf("%s", keys), ""), "")

				threeLinesXoXo[key] = true
			}
		}

	}

	return len(threeLinesXoXo)
}

func main() {
	network := parseStr(os.Args[1])

	currLength := math.MinInt
	currRes := []string{}

	for x := range network {
		cache := map[string]bool{}
		res := []string{}
		recursoveFn(network, x, cache, &res)

		if len(res) > currLength {

			slices.Sort(res)
			currRes = res
			currLength = len(res)
		}

	}

	part2 := strings.Join(currRes, ",")

	partone := part1(network)
	fmt.Println(partone, part2)
}
