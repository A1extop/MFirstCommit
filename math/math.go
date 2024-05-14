package main

import (
	"fmt"
)

func main() {
	var a int = 5
	var b int = 10
	fmt.Println(Maximum(a, b))
	fmt.Println(Minimum(a, b))
}

func Maximum[T int | int32 | int64 | float32 | float64 | uint16 | uint32 | uint64](one, two T) T {
	if one > two {
		return one
	}
	return two
}

func Minimum[T int | int32 | int64 | float32 | float64 | uint16 | uint32 | uint64](one, two T) T {
	if one < two {
		return one
	}
	return two
}
