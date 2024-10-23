package main

import "fmt"

func main() {

	a := 1.5566
	b := "yikers"

	fmt.Printf("%v %v\n", a, b)

	aSwap, bSwap := swapAnything(a, b)

	fmt.Printf("a: %v, b: %v\n", aSwap, bSwap)
}

func swapAnything(a, b any) (any, any) {
	a, b = b, a
	return a, b
}
