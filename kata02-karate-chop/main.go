package main

import (
	"fmt"
	"github.com/kejne/golang-katas/kata02-karate-chop/chop"
)

func main() {
	fmt.Printf("Found value at index %d\n",chop.Chop(3,[]int{1, 2, 3}))
}
