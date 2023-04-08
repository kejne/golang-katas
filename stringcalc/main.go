package main

import (
	"strconv"
	"strings"
)

func main() {

}

func Add(s string) int {
	nums := strings.Split(s, ",")
	sum := 0
	for _, num := range nums {

		res, err := strconv.Atoi(num)
		if err != nil {
			return 0
		}
		sum += res
	}
	return sum
}
