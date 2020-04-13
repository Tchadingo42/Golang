package main

import (
	"fmt"
)

func main() {
	var i int
	var s string
	var f float32
	var p *int
	var a []int

	p = &i
	fmt.Printf("\n int->%d\n str->%s\n float->%f\n ptr->%p\n array->%v", i, s, f, p, a)

}
