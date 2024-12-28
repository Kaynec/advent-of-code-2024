package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Network = map[string]map[string]bool

func recursoveFn(network Network, key string, cache map[string]bool, res *[]string) {
	if cache[key] || key == "" || len(network[key]) == 0 {
		return
	}

	cache[key] = true
	if !slices.Contains(*res, key) {
		*res = append(*res, key)
	}
outer:
	for newNetwork, _ := range network[key] {
		for _, oldKey := range *res {
			if !network[oldKey][newNetwork] {
				continue outer
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

	answer := strings.Join(currRes, ",")

	fmt.Println(answer)

}
