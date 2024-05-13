package main

import "fmt"

func main() {
	a := "str"
	b := "str"
	fmt.Println(Max(a, b))
}

func Max[T comparable](one T, two T) bool {
	return (one == two)
}
