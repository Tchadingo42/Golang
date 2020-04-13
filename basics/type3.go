package main

import (
	"fmt"
)

type entier int

func main() {

	var array [5]int
	for i := 0; i < 5; i++ {
		array[i] = i + 1
	}

	var slice []int
	for i := 0; i < 12; i++ {
		slice = append(slice, i+1)
	}
	fmt.Printf("\n*****Range*****\n")
	for index, x := range array {
		fmt.Println(index, x)
	}

	h := make(map[string]int)
	h["Chris"] = 42

	fmt.Printf("\n*****Array*****\n")
	fmt.Printf("%v\n", array)
	fmt.Printf("\n*****Append*****\n")
	fmt.Printf("%v\n", slice)
	fmt.Printf("\n*****Map*****\n")
	fmt.Printf("\n%v %d\n", h["Chris"])
}
