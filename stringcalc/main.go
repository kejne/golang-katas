package main

import (
	"regexp"
	"strconv"
	"strings"
)

func main() {}

func Add(s string) int {
	regSplit, err := getSplitter(s)
	if err != nil {
		return 0
	}

	sum := 0
	for _, num := range regSplit.Split(s, -1) {

		res, err := strconv.Atoi(num)
		if err != nil {
			res = 0
		}
		sum += res
	}
	return sum
}

func getSplitter(s string) (*regexp.Regexp, error) {
	del := getDelimiter(s)
	return regexp.Compile(del + "|\n")
}

func getDelimiter(s string) string {
	if len(s) < 4 {
		return ","
	}
	if !strings.Contains(s, "//") && s[:2] != "//" {
		return ","
	}
	return string(s[2])
}
